package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/blue0513/go-summarizer/httpclient"
	"github.com/blue0513/go-summarizer/openai"
	"github.com/blue0513/go-summarizer/parser"
)

func main() {
	var url, lang string
	flag.StringVar(&url, "url", "", "URL to summarize")
	flag.StringVar(&lang, "lang", "English", "Language for summarization")
	flag.Parse()

	if url == "" {
		fmt.Println("Error: URL is required")
		return
	}

	ctx := context.Background()
	html, err := httpclient.Fetch(ctx, url)
	if err != nil {
		fmt.Println("Error: parsing HTML:", err)
		return
	}

	text := parser.Extract(html)
	res, err := openai.Summarize(ctx, text, lang)
	if err != nil {
		fmt.Println("Error: summarize text", err)
		return
	}

	fmt.Println("Summarized......:\n\n", res)
}
