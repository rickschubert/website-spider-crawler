package joncalhounlinkparser

// Taken with gratitude from https://github.com/gophercises/link/blob/solution/parse.go

import (
"strings"

"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML
// document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it.
func Parse(htmlText string) []Link {
	reader := strings.NewReader(htmlText)
	doc, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links
}

func buildLink(htmlNode *html.Node) Link {
	var ret Link
	for _, attr := range htmlNode.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(htmlNode)
	return ret
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
