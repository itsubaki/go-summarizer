package httpclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

var ErrURLIsEmpty = errors.New("url is empty")

func Fetch(ctx context.Context, url string) (*html.Node, error) {
	if url == "" {
		return nil, ErrURLIsEmpty
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code=%d: %w", resp.StatusCode, err)
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse html: %w", err)
	}

	return node, nil
}
