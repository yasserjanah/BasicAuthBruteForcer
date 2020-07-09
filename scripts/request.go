package lib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

//SendHTTPRequest ... send request to the server with encoded token value
func SendHTTPRequest(url string, encodedToken string, wg *sync.WaitGroup){
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedToken))
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	resp.Body.Close()
	defer wg.Done()
	if resp.StatusCode == 200{
		fmt.Println()
		PrintSuccess("Found valid credentials", Base64Decode(encodedToken), false, false)
		fmt.Println()
		fmt.Println()
		os.Exit(0)
	}else{
		PrintFailed("failed", Base64Decode(encodedToken), false, true)
	}
}