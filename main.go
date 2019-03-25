package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func main() {
	b, err := ioutil.ReadFile(credentialsJSON())
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Set up to run 2x fetch requests to Google Sheets and GitHub concurrently
	var wg sync.WaitGroup
	var p, lf []string

	wg.Add(2)
	go func(s *sheets.Service) {
		defer wg.Done()
		p = getPatterns(s)
	}(srv)
	go func() {
		defer wg.Done()
		lf = getLintFails()
	}()
	wg.Wait()

	for _, sub := range subtractPatternsFromFails(p, lf) {
		fmt.Println(sub)
	}
}
