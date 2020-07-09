package lib

import (
	"fmt"
	"os"
)

//IsExists return whatever file is exists or Not
func IsExists(f string) bool {
    _, err := os.Stat(f)
    if os.IsNotExist(err) {
        return false
    }
    return err == nil
}

//Info return information about a given file
func Info(f string) int64 {
	fileInfo, err := os.Stat(f)
	if err != nil {
		fmt.Println("err : ", err.Error())
	}
	return fileInfo.Size()
}