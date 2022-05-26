package main

import "fmt"

func main() {

	ConfigLog()
	DeferError(xLogFile.Close)
	ConfigFlags()
	InitConfig()

	xAuthorize(
		"POST",
		xData["apiAuthorizeUrl"],
		xData["xClient"],
		xData["xSecret"],
	)

	if FlagAll {
		for cix := range Currencies {
			DisplayCard(cix)
		}
	} else {
		DisplayCard(CurrencyIx)
	}
}

func DisplayCard(cix int) {
	fmt.Printf("\nGift cards for %s %s %s %s\n",
		Currencies[cix][0], Currencies[cix][1],
		Currencies[cix][2], Currencies[cix][3])
	giftCards := PostGetGiftCards(cix)
	displayData(giftCards, cix)
}
