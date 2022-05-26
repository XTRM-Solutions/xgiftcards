package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const XtrmAPIGetGiftCards string = "/GiftCard/GetDigitalGiftCards"

func PostGetGiftCards(cix int) (r *xGetGiftCardResponse) {

	var req xGetGiftCardRequest
	var resp xGetGiftCardResponse

	req.GetGiftCards.Request.IssuerAccountNumber = xData["xIssuerID"]
	req.GetGiftCards.Request.Currency = Currencies[cix][0]
	req.GetGiftCards.Request.Pagination.RecordsToSkip = "1"
	req.GetGiftCards.Request.Pagination.RecordsToTake = strconv.Itoa(512 * 512)

	reqBuffer, err := json.Marshal(req)
	if nil != err {
		xLog.Fatalf("could not convert xGetGiftCardRequest to JSON because %s", err.Error())
	}

	respBuffer, err := XPostRequest(xData["xUrl"]+XtrmAPIGetGiftCards,
		bytes.NewReader(reqBuffer))

	if nil != err {
		xLog.Fatalf("request failed because %s\nrequest: %s",
			err.Error(), string(respBuffer[:]))
	}

	err = json.Unmarshal(respBuffer, &resp)
	if nil != err {
		if FlagDebug {
			WriteOutFile("response_buffer.debug.txt", string(respBuffer[:]))
			_, _ = fmt.Fprintf(os.Stdout, "error %s\n", err.Error())
		}
		xLog.Fatalf("Could not unmarshal response because %s",
			err.Error())
	}
	return &resp
}

type xGetGiftCardRequest struct {
	GetGiftCards struct {
		Request struct {
			IssuerAccountNumber string `json:"IssuerAccountNumber"`
			Currency            string `json:"Currency"`
			Pagination          struct {
				RecordsToSkip string `json:"RecordsToSkip"`
				RecordsToTake string `json:"RecordsToTake"`
			} `json:"Pagination"`
		} `json:"Request"`
	} `json:"GetGiftCards"`
}

func displayData(cards *xGetGiftCardResponse, cix int) {

	for _, gc := range cards.GetGiftCardResponse.GetGiftCardResult.GiftCard {
		for _, it := range gc.Items {

			switch strings.ToUpper(it.ValueType) {
			case "FIXED_VALUE":
				fmt.Printf("  SKU [ %s ]\t%s\n", it.Sku, it.RewardName)

			case "VARIABLE_VALUE":
				fmt.Printf("  SKU [ %s ]\t%s for any value between [ %s%.2f ] and [ %s%.2f ]\n",
					it.Sku,
					gc.BrandName,
					Currencies[cix][3],
					it.MinValue,
					Currencies[cix][3],
					it.MaxValue)

			default:
				xLog.Printf("Unanticipated RewardType: %s", gc.Items[0].RewardType)
			}
		}
	}
}

type xGetGiftCardResponse struct {
	GetGiftCardResponse struct {
		GetGiftCardResult struct {
			GiftCard []struct {
				BrandName   string `json:"brandName"`
				Description string `json:"description"`
				Disclaimer  string `json:"disclaimer"`
				ImageUrls   []struct {
					AdditionalProperties string `json:"additionalProperties"`
				} `json:"imageUrls"`
				Items []struct {
					Countries    []string `json:"countries"`
					CurrencyCode string   `json:"currencyCode"`
					FaceValue    float64  `json:"faceValue"`
					MaxValue     float64  `json:"maxValue"`
					MinValue     float64  `json:"minValue"`
					RewardName   string   `json:"rewardName"`
					RewardType   string   `json:"rewardType"`
					Sku          string   `json:"sku"`
					Status       string   `json:"status"`
					ValueType    string   `json:"valueType"`
				} `json:"items"`
				Terms string `json:"terms"`
			} `json:"GiftCard"`
		} `json:"GetGiftCardResult"`
		OperationStatus struct {
			Success bool   `json:"Success"`
			Errors  []byte `json:"Errors"`
		} `json:"OperationStatus"`
		PaginationTotal struct {
			RecordsToSkip int `json:"RecordsToSkip"`
			RecordsToTake int `json:"RecordsToTake"`
			RecordsTotal  int `json:"RecordsTotal"`
		} `json:"PaginationTotal"`
	} `json:"GetGiftCardResponse"`
}
