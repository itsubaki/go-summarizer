package main

import (
    "flag"
    "fmt"

    "github.com/blue0513/go-summarizer/openai"
    "github.com/blue0513/go-summarizer/parser"
    "github.com/blue0513/go-summarizer/request"
)

func main() {
    flag.Parse()

    args := flag.Args()
    if len(args) == 0 {
        fmt.Println("Error: URL is required")
        return
    }
    if len(args) > 1 {
        fmt.Println("Error: Too many arguments")
        return
    }

    url := args[0]
    doc, err := request.FetchHtml(url)
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return
    }

    text := parser.ExtractText(doc)
    response, err := openai.CallOpenAI(text)
    if err != nil {
        fmt.Println("Error calling OpenAI:", err)
        return
    }

    fmt.Println("Summarized......:\n\n", response)
}
