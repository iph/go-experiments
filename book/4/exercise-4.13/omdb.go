package main

const OmdbApi = "http://www.omdbapi.com/"

type Movie struct {
	Title     string
	Year      string
	Rated     string
	Released  string
	Genre     string
	Director  string
	Writer    string
	Actors    string
	Plot      string
	Language  string
	Country   string
	Awards    string
	Poster    string
	Metascore string
	Score     string `json:"imdbRating"`
	Votes     string `json:"imdbVotes"`
	ID        string `json:"imdbID"`
	Type      string
	Response  string
}
