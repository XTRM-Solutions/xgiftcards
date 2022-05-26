package main

import (
	"bufio"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/spf13/pflag"
	"io"
	"log"
	"os"
	"strings"
)

var xData = map[string]string{
	"currentSection": ini.DefaultSection,
}

const DefaultCurrencyCode = "EUR"
const ProgramName = "XtrmGiftCards"
const LogName = ProgramName + ".log"

/*
const RegExpAccountPattern = "^(PAT|SPN)[0-9]+$"
const RegExpTaxYearPattern = "^2[012][0-9]$"
*/

var xLogFile *os.File
var xLogBuffer *bufio.Writer
var xLog log.Logger
var FlagAll bool
var FlagDebug bool

/* var FlagQuiet bool */

var FlagVerbose bool
var CurrencyIx int
var nFlags *pflag.FlagSet

func ConfigLog() {

	var err error
	var logWriters []io.Writer

	xLogFile, err = os.OpenFile(ProgramName+".log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if nil != err {
		xLog.Fatalf("error opening log file %s: %v", LogName, err)
	}

	xLogBuffer = bufio.NewWriter(xLogFile)

	logWriters = append(logWriters, os.Stderr)
	logWriters = append(logWriters, xLogBuffer)

	xLog.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
	xLog.SetOutput(io.MultiWriter(logWriters...))
}

func ConfigFlags() {
	nFlags = pflag.NewFlagSet("default", pflag.ContinueOnError)

	nFlags.StringP("currency", "c", DefaultCurrencyCode, "Three-Letter Currency Code")
	nFlags.BoolP("debug", "d", false, "Enable additional debugging info")
	nFlags.BoolP("quiet", "q", false, "Do not log to stdout, only log file")
	nFlags.BoolP("all", "a", false, "Show all cards for all currencies")
	nFlags.StringP("profile", "p", "DEFAULT", "Select profile within INI file")
	nFlags.BoolP("help", "h", false, "print usage message")

	err := nFlags.Parse(os.Args[1:])
	if nil != err {
		xLog.Fatalf("\nerror parsing flags: %s\n%s %s\n%s\n\t%v\n",
			err.Error(),
			"common issue: 2 hyphens for long-form arguments,",
			"1 hyphen for short-form argument",
			"Program arguments are:",
			os.Args)
	}

	if GetFlagBool("help") {
		UsageMessage()
		os.Exit(0)
	}

	if GetFlagBool("all") {
		FlagAll = true
	} else {
		currencyCode := strings.ToUpper(GetFlagString("currency"))
		CurrencyIx = -1
		for ix, val := range Currencies {
			if strings.EqualFold(currencyCode, val[0]) {
				CurrencyIx = ix
				break
			}
		}
		if CurrencyIx < 0 {
			UsageMessage()
			xLog.Printf("\n\tOops! Invalid currency %s selected", currencyCode)
			xLog.Fatalf("\n\tPlease use a valid currency code\n. Use -h or --help for help.\n")
		}

	}

	if GetFlagBool("quiet") {
		// FlagQuiet = true
		xLog.SetOutput(xLogBuffer)
	} /* else {
		FlagQuiet = false
	} */

}

// GetFlagBool fetch the bool for a boolean flag
func GetFlagBool(key string) (value bool) {
	var err error
	value, err = nFlags.GetBool(key)
	if nil != err {
		xLog.Fatalf("error fetching value for boolean flag [ %s ]: %s \n", key, err.Error())
		return false
	}
	return value
}

// GetFlagString fetch the string associated with a CLI arg
func GetFlagString(key string) (value string) {
	var err error
	value, err = nFlags.GetString(key)
	if nil != err {
		xLog.Fatalf("error fetching value for string flag [ %s ]: %s \n", key, err.Error())
		return ""
	}
	return value
}

/*
// GetFlagInt fetch the value of integer flag
func GetFlagInt(key string) (value int) {
	var err error
	value, err = nFlags.GetInt(key)
	if nil != err {
		xLog.Fatalf("%s [ %s ]: %s \n",
			"error fetching value for integer flag",
			key, err.Error())
		return 0
	}
	return value
}
*/

// UsageMessage
// print message about program & arguments
func UsageMessage() {

	_, _ = fmt.Printf("%s: list gift cards for selected currency (default: %s)\n",
		ProgramName, DefaultCurrencyCode)
	_, _ = fmt.Printf("Accepted currency codes are:\n")
	for _, val := range Currencies {
		_, _ = fmt.Printf("\t%s %s (currency of %s, symbol[ %s ])\n", val[0], val[2], val[1], val[3])
	}
	_, _ = fmt.Printf("\nAs some currencies share full names, please use the three-letter currency code.\n\n")
	_, _ = fmt.Printf("Program Usage: \n")
	nFlags.PrintDefaults()

}

// Currencies
// please note that this [][]string is immutable,
// it's an array not a slice
// ==> insertions should be done in alpha order for my sanity!
// Format is currency code, nation, name of currency, and the symbol or abbreviation for the currency
// could hold all this in the ini file
var Currencies = [...][4]string{
	{"AED", "United Arab Emirates", "Dirham", "د.إ"},
	{"AUD", "Australia", "Dollar", "AU$"},
	{"BGN", "Bulgaria", "Lev", "лв"},
	{"BHD", "Bahrain", "Dinar", ".د.ب"},
	{"CAD", "Canada", "Dollar", "CAD$"},
	{"CHF", "Switzerland", "Franc", "SFr"},
	{"CNH", "China", "Yuan", "¥"},
	{"CZK", "Czech Republic", "Koruna", "Kč"},
	{"DKK", "Denmark", "Krone", "DKK kr"},
	{"EUR", "EU", "Euro", "€"},
	{"FJD", "Fiji", "Dollar", "FJ$"},
	{"GBP", "UK", "Pound", "GB£"},
	{"HKD", "Hong Kong", "Dollar", "HK$"},
	{"HRK", "Croatia", "Kuna", "kn"},
	{"HUF", "Hungary", "Forint", "Ft"},
	{"IDR", "Indonesia", "Rupiah", "Rp"},
	{"ILS", "Israel", "Shekel", "₪"},
	{"JOD", "Jordan", "Dinar", "د.أ"},
	{"JPY", "Japan", "Yen", "¥"},
	{"KWD", "Kuwait", "Dinar", "د.ك"},
	{"MAD", "Morocco", "Dirham", "د.م."},
	{"NOK", "Norway", "Krone", "NOK kr"},
	{"MUR", "Mauritian", "Rupee", "₨"},
	{"MXN", "Mexico", "Peso", "M$"},
	{"NZD", "New Zealand", "Dollar", "NZ$"},
	{"OMR", "Oman", "Rial", "ر.ع."},
	{"PLN", "Poland", "Zloty", "zł"},
	{"QAR", "Qatar", "Riyal", "ر.ق"},
	{"RON", "Romania", "Leu", "RL"},
	{"SEK", "Sweden", "Krona", "SEK"},
	{"SGD", "Singapore", "Dollar", "S$"},
	{"THB", "Thailand", "Baht", "฿"},
	{"TND", "Tunisia", "Dinar", "د.ت"},
	{"TRY", "Turkey", "Lira", "₺"},
	{"USD", "United States", "Dollar", "US$"},
	{"ZAR", "South Africa", "Rand", "R"},
}
