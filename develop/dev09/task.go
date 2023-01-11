package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	file, err := download(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file with name: %s was created\n", file.Name())
}

func download(path string) (*os.File, error) {
	if err := parseURL(path); err != nil {
		return nil, err
	}
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	file, err := os.Create("index.html")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if _, err := file.Write(body); err != nil {
		return nil, err
	}
	return file, nil
}

func parseURL(rawURL string) error {
	_, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	return nil
}
