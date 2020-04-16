package utils

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func findTitle(r io.Reader) (string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", err
	}

	var f func(*html.Node) (string, bool)

	f = func(node *html.Node) (string, bool) {
		if node.Type == html.ElementNode && node.Data == "title" {
			return node.FirstChild.Data, true
		}

		for ch := node.FirstChild; ch != nil; ch = ch.NextSibling {
			if title, ok := f(ch); ok {
				return title, true
			}

		}

		return "", false
	}
	title, _ := f(doc)
	return title, nil
}

// FetchTitle finds the title of a page.
func FetchTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	title := "Untitled"

	contentType := resp.Header.Values("Content-Type")
	if len(contentType) > 0 && strings.HasPrefix(contentType[0], "text/html") {
		buf := make([]byte, 16*1024)

		n, err := io.ReadFull(resp.Body, buf)
		buf = buf[:n]

		br := bytes.NewReader(buf)
		t, err := findTitle(br)
		if err != nil {
			return "", err
		}

		t = strings.Trim(t, " \r\n\t")
		if t != "" {
			title = t
		}
	}

	return title, nil
}
