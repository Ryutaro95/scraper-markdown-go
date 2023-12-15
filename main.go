package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/chromedp/chromedp"
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

func getHTML(url string) (string, string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var res string
	var title string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("article"),
		chromedp.OuterHTML("article", &res, chromedp.ByQuery),
		chromedp.Text("h1", &title, chromedp.ByQuery),
	)
	if err != nil {
		return "", "", err
	}
	return res, title, nil
}

func htmlToMarkdown(html string) (string, error) {
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(html)
	if err != nil {
		return "", err
	}
	return markdown, nil
}

func saveMarkdowntoFile(markdown, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(markdown)
	if err != nil {
		return err
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

	html, title, err := getHTML(url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(title)
	markdown, err := htmlToMarkdown(html)
	if err != nil {
		log.Fatalf("err html to markdown: %s\n", err)
	}

	filename := fmt.Sprintf("%s.md", title)
	if err := saveMarkdowntoFile(markdown, filename); err != nil {
		log.Fatalf("failed save markdown: %s\n", err)
	}
	fmt.Printf("markdown content has saved to %s\n", filename)
}
