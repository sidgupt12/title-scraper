package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetBlogTitles(url, selector string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	titles := ""
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		titles += "-" + s.Text() + "\n"

	})
	return titles, nil

}
