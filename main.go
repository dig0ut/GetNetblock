package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

//add your API key here (free from https://ip-netblocks.whoisxmlapi.com/api/signup)
var apiKey = ""

func searchIP(IP string) {

	white := color.New(color.FgWhite)
	bold := white.Add(color.Bold)

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
	asName := gjson.Get(string(Data), "result.inetnums.0.as.name")
	inetnums := gjson.Get(string(Data), "result.inetnums.0.inetnum")
	route := gjson.Get(string(Data), "result.inetnums.0.as.route")
	searched := gjson.Get(string(Data), "search")
	bold.Println("Results for", searched.String()+":")
	fmt.Println("ASN is owned by: ", asName.String())
	fmt.Println("Netblock for IP: ", inetnums.String())
	fmt.Println("CIDR equivelant: ", route.String(), "\n")

}

func searchOrg(Org string) {

	white := color.New(color.FgWhite)
	bold := white.Add(color.Bold)

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
	bold.Println("CIDR ranges associated with", Org+":")
	res.ForEach(func(key, value gjson.Result) bool {
		println(value.String())
		return true
	})

}

func main() {

	ipAddress := flag.String("ip", "", "IP address.")
	orgName := flag.String("org", "", "Organisation name.")
	sourceFile := flag.String("source", "", "File containing IP's to query (one per line).")
	flag.Parse()

	// Check for API key
	if len(apiKey) == 0 {
		fmt.Println("You need to include your API key in the 'apiKey' variable within main.go, then recompile and try again.")
		os.Exit(1)
	}

	// If source file is specified, open file, convert to array and query each IP
	if len(*sourceFile) > 0 {

		file, err := os.Open(*sourceFile)
		if err != nil {
			log.Fatalf("open file error: %v", err)
			return
		}
		//i := 0
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var entry []string
		lineCount := 0
		for scanner.Scan() {
			lineCount++
			entry = append(entry, scanner.Text())
		}
		fmt.Println("Total of", lineCount, "IP's received from", *sourceFile)
		fmt.Println("Querying...\n")
		file.Close()

		for _, each_ln := range entry {
			searchIP(each_ln)
		}

	}

	// Require -ip arg to be given
	if len(*ipAddress) == 0 {
		if len(*orgName) == 0 {
			if len(*sourceFile) == 0 {
				fmt.Println("No IP address or org name specified.")
				fmt.Println("Usage:")
				flag.PrintDefaults()
				os.Exit(1)
			}
		}
	}

	if len(*ipAddress) > 0 {
		searchIP(*ipAddress)
	}

	if len(*orgName) > 0 {
		searchOrg(*orgName)
	}

}
