package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// FetchMostStarredRepositories is a function that fetches most starred repositories on GitHub
func FetchMostStarredRepositories() *RepositoriesSearchResponse {
	response := new(RepositoriesSearchResponse)
	getJSON("https://api.github.com/search/repositories?q=stars:%3E1000", response)
	return response
}
