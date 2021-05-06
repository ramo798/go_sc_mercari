package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func get_details_item(url string) (bool, string, string, string) {
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
	}

	var soldout bool
	sold := doc.Find(".item-box-container .item-sold-out-badge")
	if sold.Text() == "" {
		soldout = false
	} else {
		soldout = true
	}

	title := doc.Find(".item-name").First().Text()

	return soldout, no, title, "https://www.mercari.com" + url
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	url := "/jp/u/951762445/"

	page_url := get_all_page_url(url)
	fmt.Println(page_url)

	items_url := get_items_url(page_url)
	fmt.Println(len(items_url))

	file, err := os.OpenFile("mercari_list.csv", os.O_WRONLY|os.O_CREATE, 0600)
	failOnError(err)
	defer file.Close()
	err = file.Truncate(0)
	failOnError(err)
	writer := csv.NewWriter(file)

	for _, url := range items_url {
		soldout, no, title, item_url := get_details_item(url)
		fmt.Println(soldout, no, title, item_url)
		writer.Write([]string{strconv.FormatBool(soldout), no, title, item_url})
	}

	writer.Flush()

}
