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

func searchIP(IP string) {

	// Build the GET request
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
	route := gjson.Get(string(Data), "result.inetnums.0.as.route")
	//country := gjson.Get(string(Data), "result.inetnums.0.country")
	fmt.Println("ASN is owned by: ", asName.String())
	fmt.Println("Netblock for IP: ", inetnums.String())
	fmt.Println("CIDR equivelant: ", route.String())
	//fmt.Println("Originates from: ", country.String())

}

func searchOrg(Org string) {

	// Build the GET request
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ip-netblocks.whoisxmlapi.com/api/v2", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// Paramaters to be included in the GET request
	params := req.URL.Query()
	params.Add("apiKey", apiKey)
	params.Add("org", Org)
	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Convert response to string and parse required json values
	Data, _ := ioutil.ReadAll(resp.Body)
	res := gjson.Get(string(Data), "result.inetnums.#.as.route")
	fmt.Println("CIDR ranges associated with", Org, ":")
	res.ForEach(func(key, value gjson.Result) bool {
		println(value.String())
		return true
	})

}

func main() {

	// Check for API key
	if len(apiKey) == 0 {
		fmt.Println("You need to include your API key in the 'apiKey' variable within main.go, then recompile and try again.")
		os.Exit(1)
	}

	ipAddress := flag.String("ip", "", "IP address.")
	orgName := flag.String("org", "", "Organisation name.")
	flag.Parse()

	// Require -ip arg to be given
	if len(*ipAddress) == 0 {
		if len(*orgName) == 0 {
			fmt.Println("No IP address or org name specified.")
			fmt.Println("Usage:")
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	if len(*ipAddress) > 0 {
		searchIP(*ipAddress)
	}

	if len(*orgName) > 0 {
		searchOrg(*orgName)
	}

}
