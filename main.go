package main

import (
    "flag"
    "fmt"

    "github.com/blue0513/go-summarizer/openai"
    "github.com/blue0513/go-summarizer/parser"
    "github.com/blue0513/go-summarizer/request"
)

func main() {
    url := flag.String("url", "", "URL to summarize")
    lang := flag.String("lang", "English", "Language for summarization")
    flag.Parse()

    if url == nil || *url == "" {
        fmt.Println("Error: URL is required")
        return
    }

    doc, err := request.FetchHtml(*url)
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return
    }

    text := parser.ExtractText(doc)
    response, err := openai.CallOpenAI(text, *lang)
    if err != nil {
        fmt.Println("Error calling OpenAI:", err)
        return
    }

    fmt.Println("Summarized......:\n\n", response)
}
