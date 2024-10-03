package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	url := flag.String("url", "", "The URL of the website to scrape")
	selector := flag.String("selector", "", "The CSS selector for the blog titles")
	flag.Parse()

	if *url == "" || *selector == "" {
		log.Fatal("Both url and selector flags are required")
	}
	blogTitles, err := GetBlogTitles(*url, *selector)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Blog Titles: ")
	fmt.Println(blogTitles)

}
