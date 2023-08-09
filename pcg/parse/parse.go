package parse

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func Parse() []string {
	url := "https://go-proverbs.github.io/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var proverbs []string
	var extractProverbs func(*html.Node)
	extractProverbs = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" && n.Parent != nil && n.Parent.Data == "h3" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					proverbs = append(proverbs, c.Data)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractProverbs(c)
		}
	}

	extractProverbs(doc)
	return proverbs

}
