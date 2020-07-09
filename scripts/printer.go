package lib

import (
	"github.com/fatih/color"
	//"os"
	"fmt"
)

//inSameLine : a string to clear the line to print on it again
const inSameLine string = "\033[G\033[K"

//PrintSuccess : print a success message
func PrintSuccess(pr, text string, colored bool, newline bool){
	fmt.Print(inSameLine)
	if newline{
		text = text + ".\n"
	} else {
		text = text + "."
	}
	BOLDFg := color.New(color.FgWhite, color.Bold).SprintFunc()
	GREENBoldFg := color.New(color.FgGreen, color.Bold).SprintFunc()
	if colored{
		FgGreen := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s %s: %s", GREENBoldFg("[+]"), BOLDFg(pr), FgGreen(text))
		return 
	}
	fmt.Printf("%s %s: %s", GREENBoldFg("[+]"), BOLDFg(pr), text)
}

//PrintFailed : print a failed message
func PrintFailed(pr, text string, colored bool, newline bool){
	fmt.Print(inSameLine)
	BOLDFg := color.New(color.FgYellow, color.Bold).SprintFunc()
	REDBoldFg := color.New(color.FgRed, color.Bold).SprintFunc()
	if colored{
		FgRed := color.New(color.FgRed).SprintFunc()
		if newline{
			fmt.Printf("%s %s %s%s%s.\n", REDBoldFg("[-]"), BOLDFg(pr), BOLDFg("'"), FgRed(text), BOLDFg("'"))
			return
		}
		fmt.Printf("%s %s %s%s%s.", REDBoldFg("[-]"), BOLDFg(pr), BOLDFg("'"), FgRed(text), BOLDFg("'"))
		return 
	}
	if newline{
		fmt.Printf("%s %s %s%s%s.\n", REDBoldFg("[-]"), BOLDFg(pr), BOLDFg("'"), text, BOLDFg("'"))
		return
	}
	fmt.Printf("%s %s %s%s%s.", REDBoldFg("[-]"), BOLDFg(pr), BOLDFg("'"), text, BOLDFg("'"))
}

//PrintStatus : print a status message
func PrintStatus(pr, text string, colored bool, newline bool){
	fmt.Print(inSameLine)
	if newline{
		text = text + ".\n"
	} else {
		text = text + "."
	}
	BOLDFg := color.New(color.FgWhite, color.Bold).SprintFunc()
	BLUEBoldFg := color.New(color.FgBlue, color.Bold).SprintFunc()
	if colored{
		FgBlue := color.New(color.FgBlue).SprintFunc()
		fmt.Printf("%s %s: %s", BLUEBoldFg("[*]"), BOLDFg(pr), FgBlue(text))
		return 
	}
	fmt.Printf("%s %s: %s", BLUEBoldFg("[*]"), BOLDFg(pr), text)
}