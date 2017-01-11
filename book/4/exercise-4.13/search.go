package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func SearchMovie(movieName string) (*Movie, error) {
	q := url.QueryEscape(movieName)
	resp, err := http.Get(OmdbApi + "?t=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//!-
