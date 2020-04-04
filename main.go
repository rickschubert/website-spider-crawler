package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/rickschubert/website-spider-crawler/joncalhounlinkparser"
)

var linksVisited []joncalhounlinkparser.Link

func main() {
	fmt.Println("The website crawler has started!")
	baseUrl := getBaseUrlFromCLIFlags()
	fmt.Println("baseUrl", baseUrl)
	htmlContent := getHTMLContent(baseUrl)
	linksFromHTML := joncalhounlinkparser.Parse(htmlContent)
	fmt.Println(linksFromHTML)
}

func getBaseUrlFromCLIFlags() string {
	parameterName := "baseUrl"
	baseUrlPtr := flag.String(parameterName, "", "The website from which you want to start performing the spider crawl. I.e. \"https://starwars.com\"")
	flag.Parse()
	exitIfNoBaseUrlWasProvided(baseUrlPtr, parameterName)
	return *baseUrlPtr
}

func exitIfNoBaseUrlWasProvided(baseUrlPtr *string, parameterName string) {
	if *baseUrlPtr == "" {
		log.Fatal(fmt.Sprintf("You need to specify the parameter \"%s\". Start the problem with the flag -h for reference of how to do this.", parameterName))
	}
}

func getHTMLContent(url string) string {
	response, err := http.Get(url)
	panicError(fmt.Sprintf("Error: Unable to get HTTP response from url %s", url), err)
	body, err := ioutil.ReadAll(response.Body)
	panicError(fmt.Sprintf("Error: Unable to read response body from URL %s", url), err)
	err = response.Body.Close()
	panicError(fmt.Sprintf("Error: Unable to close HTTP conneection made to url %s", url), err)
	return string(body)
}

func panicError(errorMessage string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s ---> %s", errorMessage, err))
	}
}

func trimToFirstOneHundredCharacters(str string) string {
	return str[0:99]
}
