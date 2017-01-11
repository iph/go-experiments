package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	result, err := SearchMovie(os.Args[1])
	if err != nil || result.Response == "False" {
		log.Fatal(err)
	}

	posterImage, err := http.Get(result.Poster)
	defer posterImage.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	extension := strings.Split(result.Poster, ".")
	file, _ := os.Create("image." + extension[len(extension)-1])
	defer file.Close()

	io.Copy(file, posterImage.Body)
}
