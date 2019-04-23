package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"google.golang.org/api/sheets/v4"
)

func getPatterns(srv *sheets.Service) []string {
	// Collects the patterns listed on the Google Sheet into a slice of strings
	// https://docs.google.com/spreadsheets/d/1VhU6zCk-vaAnbu_XkgsIq5I7jVnezUquWBMhcIN6AJM/edit#gid=0
	spreadsheetID := "1VhU6zCk-vaAnbu_XkgsIq5I7jVnezUquWBMhcIN6AJM"
	readRange := "Sheet1!A2:A99999"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	patterns := make([]string, 0)

	if len(resp.Values) > 0 {
		for _, b := range resp.Values {
			if len(b) > 0 && len(b[0].(string)) > 0 {
				patterns = append(patterns, b[0].(string))
			}
		}
	}

	return patterns
}

func getLintFails() []string {
	raw := "https://github.com/kubernetes/kubernetes/raw/master/hack/.golint_failures"
	req, errReq := http.NewRequest(http.MethodGet, raw, nil)
	if errReq != nil {
		log.Fatalf("Error creating HTTP request: %v", errReq)
	}

	resp, errResp := http.DefaultClient.Do(req)
	if errResp != nil {
		log.Fatalf("Error from HTTP response: %v", errResp)
	}

	defer resp.Body.Close()
	body, errBodyRead := ioutil.ReadAll(resp.Body)
	if errBodyRead != nil {
		log.Fatalf("Error reading response body: %v", errBodyRead)
	}

	fails := make([]string, 0)

	for _, line := range strings.Split(string(body), "\n") {
		if len(line) > 0 {
			fails = append(fails, line)
		}
	}

	return fails
}
