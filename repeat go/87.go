package main

import (
	"html"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == 'a' {
		for _, a := range n.Attr {
			if a.key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links

}
