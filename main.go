package main

import (
	"flag"
	"fmt"
	"log"
	//"github.com/rickschubert/html-link-parser/htmllinkextractor"
)

func main() {
	fmt.Println("The website crawler has started!")
	baseUrl := getBaseUrlFromCLIFlags()
	fmt.Println("baseUrl", baseUrl)
	//extractedStuff := htmllinkextractor.ExtractLinks("")
}

func getBaseUrlFromCLIFlags() string {
	parameterName := "baseUrl"
	baseUrlPtr := flag.String(parameterName, "", "The website from which you want to start performing the spider crawl. I.e. \"https://starwars.com\"")
	flag.Parse()
	exitIfNoBaseUrlWasProvided(baseUrlPtr, parameterName)
	return ""
}

func exitIfNoBaseUrlWasProvided(baseUrlPtr *string, parameterName string) {
	if *baseUrlPtr == "" {
		log.Fatal(fmt.Sprintf("You need to specify the parameter \"%s\". Start the problem with the flag -h for reference of how to do this.", parameterName))
	}
}
