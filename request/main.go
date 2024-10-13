package request

import (
    "fmt"
    "net/http"

    "golang.org/x/net/html"
)

func FetchHtml(url string) (*html.Node, error) {
    if url == "" {
        err := fmt.Errorf("Error: URL is required")
        return nil, err
    }

    resp, err := http.Get(url)
    if err != nil {
        err := fmt.Errorf("Error fetching URL: %v", err)
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        err := fmt.Errorf("Error: Status code: %d", resp.StatusCode)
        return nil, err
    }

    return html.Parse(resp.Body)
}
