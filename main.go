package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/imroc/req"
)

func main() {
	token := os.Getenv("PLUGIN_TOKEN")
	botname := os.Getenv("PLUGIN_BOTNAME")
	message := os.Getenv("PLUGIN_MESSAGE")
	header := make(http.Header)
	header.Set("Accept", "application/json")
	r, err := req.Post(fmt.Sprintf("https://cliq.zoho.com/api/v2/bots/%s/incoming?zapikey=%s", botname, token), header, message)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	if r.Response().StatusCode != http.StatusOK {
		fmt.Println("HTTP Status", r.Response().StatusCode)
		os.Exit(1)
	}
	fmt.Println("Success")
	fmt.Println(token)
	fmt.Println(botname)
	fmt.Println(message)
	os.Exit(0)
}
