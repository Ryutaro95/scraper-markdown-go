package main

import (
	"flag"
	"fmt"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "Specify the URL for the tool to interact with.(long form)")
	flag.StringVar(&url, "u", "", "Specify the URL for the tool to interact with.")
	flag.Parse()
	fmt.Println(url)
}
