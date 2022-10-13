package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

//add your API key here (free from https://ip-netblocks.whoisxmlapi.com/api/signup)
var apiKey = ""

func httpGetNetblocks(IP string) {

	// Build the request
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ip-netblocks.whoisxmlapi.com/api/v2", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// Paramaters to be included in the GET request
	params := req.URL.Query()
	params.Add("apiKey", apiKey)
	params.Add("ip", IP)
	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Convert response to string and parse required json values 
	Data, _ := ioutil.ReadAll(resp.Body)
	inetnums := gjson.Get(string(Data), "result.inetnums.0.inetnum")
	asName := gjson.Get(string(Data), "result.inetnums.0.as.name")
	fmt.Println("ASN is owned by: ", asName.String())
	fmt.Println("Netblock for IP: ", inetnums.String())

}

func main() {

	ipAddress := flag.String("ip", "", "IP address.")
	flag.Parse()

	httpGetNetblocks(*ipAddress)

}
