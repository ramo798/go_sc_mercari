package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func get_items_url(urls []string) []string {
	var resUrl []string

	for _, url := range urls {
		res, err := http.Get("https://www.mercari.com" + url)
		if err != nil {
			log.Println(err)
		}
		defer res.Body.Close()

		doc, _ := goquery.NewDocumentFromReader(res.Body)
		doc.Find(".items-box a").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			resUrl = append(resUrl, url)
			// fmt.Println(url)
		})
	}

	return resUrl
}

func get_all_page_url(origin_url string) []string {
	url_list := []string{origin_url}

	res, err := http.Get("https://www.mercari.com" + origin_url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	next_url, exist := doc.Find(".pager-next .pager-cell a").Attr("href")

	if exist {
		url_list = append(url_list, next_url)
		for exist {
			res, err := http.Get("https://www.mercari.com" + next_url)
			if err != nil {
				log.Println(err)
			}
			defer res.Body.Close()
			doc, _ := goquery.NewDocumentFromReader(res.Body)
			next_url, exist = doc.Find(".pager-next .pager-cell a").Attr("href")
			url_list = append(url_list, next_url)
		}
	}

	return url_list
}

func get_details_item(url string) (string, string, string) {
	res, err := http.Get("https://www.mercari.com" + url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	text := doc.Find(".item-description p").First().Text()
	var no string
	if strings.Contains(text, "[") && strings.Contains(text, "]") {
		no = text[strings.Index(text, "[")+1 : strings.Index(text, "]")]
		// fmt.Print(text[strings.Index(text, "[")+1 : strings.Index(text, "]")])
	}
	title := doc.Find(".item-name").First().Text()
	// fmt.Println(text)
	// fmt.Println(title)
	return no, title, "https://www.mercari.com" + url
}

func main() {
	// url := "/jp/u/951762445/"

	// page_url := get_all_page_url(url)
	// fmt.Println(page_url)

	// items_url := get_items_url(page_url)
	// fmt.Println(len(items_url))

	no, title, url := get_details_item("/jp/items/m46693656082/")
	fmt.Println(no, title, url)
}
