package main

import (
	"errors"
	"flag"
	"log"
	"net/url"
)

func validateURL(urlStr string) error {
	if urlStr == "" {
		return errors.New("url is required")
	}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return err
	}

	if parsedURL.Scheme == "" {
		return errors.New("misisng url scheme")
	}

	if parsedURL.Host == "" {
		return errors.New("missing url host")
	}

	return nil
}

func main() {
	var url string
	flag.StringVar(&url, "url", "", "Specify the URL for the tool to interact with.(long form)")
	flag.StringVar(&url, "u", "", "Specify the URL for the tool to interact with.")
	flag.Parse()

	if err := validateURL(url); err != nil {
		log.Fatalf("Invalid URL: %s\n", err)
	}
}
