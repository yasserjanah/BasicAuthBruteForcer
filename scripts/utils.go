package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	b64 "encoding/base64"
	"github.com/fatih/color"
)

//CheckFlag : Validate flag for errors
func CheckFlag(s string) (string, error) {
	if (s == ""){
		return "", errors.New("is required")
	}
	return s, nil
}

// Usage : print usage
func Usage(){
	BOLDFg := color.New(color.FgWhite, color.Bold).SprintFunc()
	YellowBOLDFg := color.New(color.FgYellow, color.Bold).SprintFunc()
	CyanBoldFg := color.New(color.FgCyan, color.Bold).SprintFunc()
	MagBoldFg := color.New(color.FgMagenta, color.Bold).SprintFunc()
	fmt.Fprintf(os.Stdout,"%s %s.\n", BOLDFg("Usage of"), CyanBoldFg(os.Args[0]))
	fmt.Fprintf(os.Stdout,"\t%s %s.\n", YellowBOLDFg("-user"), MagBoldFg("string"))
	fmt.Fprintf(os.Stdout,"\t\t One User to test.\n")
	fmt.Fprintf(os.Stdout,"\t%s %s.\n", YellowBOLDFg("-passwd"), MagBoldFg("string"))
	fmt.Fprintf(os.Stdout,"\t\t One Password to test.\n")
	fmt.Fprintf(os.Stdout,"\t%s %s.\n", YellowBOLDFg("-user-file"), MagBoldFg("FILE"))
	fmt.Fprintf(os.Stdout,"\t\t List of users.\n")
	fmt.Fprintf(os.Stdout,"\t%s %s.\n", YellowBOLDFg("-passwd-file"), MagBoldFg("FILE"))
	fmt.Fprintf(os.Stdout,"\t\t List of passwords.\n")
	os.Exit(0)
}


//ReadFile ... read wordlist and return a slice 
func ReadFile(f string) []string{
	data := []string{}
	fi, err := os.Open(f)
	if err != nil {
		fmt.Println("error: cannot read file.")
		os.Exit(0)
	}
	var file *bufio.Scanner = bufio.NewScanner(fi)
	for file.Scan(){
		if file.Text() != ""{
			data = append(data, strings.TrimRight(file.Text(), "\r\n"))	
		}
	}
	return data
}

//Base64Encode ... encode string to base64
func Base64Encode(str string) string{
	sEnc := b64.StdEncoding.EncodeToString([]byte(str))
	return string(sEnc)

}

//Base64Decode ... decode base64 to string
func Base64Decode(sEnc string) string{
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	return string(sDec)
}


//SplitSlice ... divide a slice into 10 pieces
func SplitSlice(s []string, numCPU int) [][]string{
	var divided [][]string
	chunkSize := (len(s) + numCPU - 1) / numCPU

	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize

		if end > len(s) {
			end = len(s)
		}

		divided = append(divided, s[i:end])
	}

	return divided
}