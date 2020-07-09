package lib

import (
	"math"
	"strconv"
	"github.com/fatih/color"
)

//RoundUp ...
func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}

//ByteFormat ...
func ByteFormat(inputNum float64, precision int) string {
	BOLDWhite := color.New(color.FgWhite, color.Bold).SprintFunc()
	GREEN := color.New(color.FgGreen).SprintFunc()

	if precision <= 0 {
			precision = 1
	}

	var unit string
	var returnVal float64

	if inputNum >= 1000000000000000000000000 {
			returnVal = RoundUp((inputNum / 1208925819614629174706176), precision)
			unit = BOLDWhite(" YB") // yottabyte
	} else if inputNum >= 1000000000000000000000 {
			returnVal = RoundUp((inputNum / 1180591620717411303424), precision)
			unit = BOLDWhite(" ZB") // zettabyte
	} else if inputNum >= 10000000000000000000 {
			returnVal = RoundUp((inputNum / 1152921504606846976), precision)
			unit = BOLDWhite(" EB") // exabyte
	} else if inputNum >= 1000000000000000 {
			returnVal = RoundUp((inputNum / 1125899906842624), precision)
			unit = BOLDWhite(" PB") // petabyte
	} else if inputNum >= 1000000000000 {
			returnVal = RoundUp((inputNum / 1099511627776), precision)
			unit = BOLDWhite(" TB") // terrabyte
	} else if inputNum >= 1000000000 {
			returnVal = RoundUp((inputNum / 1073741824), precision)
			unit = BOLDWhite(" GB") // gigabyte
	} else if inputNum >= 1000000 {
			returnVal = RoundUp((inputNum / 1048576), precision)
			unit = BOLDWhite(" MB") // megabyte
	} else if inputNum >= 1000 {
			returnVal = RoundUp((inputNum / 1024), precision)
			unit = BOLDWhite(" KB") // kilobyte
	} else {
			returnVal = inputNum
			unit = BOLDWhite(" bytes") // byte
	}

	return GREEN(strconv.FormatFloat(returnVal, 'f', precision, 64)) + unit

}