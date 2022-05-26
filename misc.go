package main

import (
	"fmt"
	"os"
	"strings"
)

// DeferError
// accounts for an at-close function that
// may return an error for its close
func DeferError(f func() error) {
	err := f()
	if nil != err {
		xLog.Printf("%s%s",
			"(may be harmless) error in deferred function: ",
			err.Error())
	}
}

// WriteSB Add a series of strings to a strings.Builder
func WriteSB(sb *strings.Builder, inputStrings ...string) {
	for _, val := range inputStrings {
		_, err := sb.WriteString(val)
		if nil != err {
			xLog.Print("strings.Builder failed to add " + val + " ??")
			xLog.Fatal("values: ", inputStrings)
		}
	}
}

func IsStringSet(stringToCheck *string) (isSet bool) {
	if nil != stringToCheck && "" != *stringToCheck {
		return true
	}
	return false
}

// WriteOutFile Write out something into a file & close it, no buffering, for debug purposes
func WriteOutFile(fullFileName, log string) {
	var debugFile *os.File = nil
	var err error = nil

	if !IsStringSet(&fullFileName) || !IsStringSet(&log) {
		_, _ = fmt.Fprintf(os.Stderr,
			"tried to write a debug file but missing filename or logdata or both\n")
	}

	debugFile, err = os.OpenFile(fullFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if nil != err {
		_, _ = fmt.Fprintf(os.Stdout, "could not open scratch file %s because %s\n",
			fullFileName, err.Error())
		os.Exit(2)
	}
	DeferError(debugFile.Close)

	_, err = debugFile.WriteString(log)
	if nil != err {
		_, _ = fmt.Fprintf(os.Stdout, "could not not write to scratch file %s because %s\n",
			fullFileName, err.Error())
		os.Exit(3)
	}

}

func CheckErr(err error) {
	if nil != err {
		xLog.Panicf("Some functions return errors, but errors are so \n"+
			"unusual that it's not worth handling them ... but we can still\n"+
			"panic for a stack trace. Hint: *that just happened!*\n"+
			"Error text: %s \n", err.Error())
	}
}
