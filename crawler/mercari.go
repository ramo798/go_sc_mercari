package crawler

import (
	"log"
	"net/http"
	"strings"
	"test/model"

	"github.com/PuerkitoBio/goquery"
)

func Get_items_url(url string) []string {
	var resUrl []string

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".items-box a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		resUrl = append(resUrl, "https://www.mercari.com"+url)
		// fmt.Println(url)
	})

	return resUrl
}

func Get_all_page_url(uid string) []string {
	origin_url := "https://www.mercari.com/jp/u/" + uid + "/"

	url_list := []string{origin_url}

	res, err := http.Get(origin_url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	next_url, exist := doc.Find(".pager-next .pager-cell a").Attr("href")

	if exist {
		url_list = append(url_list, "https://www.mercari.com"+next_url)
		for exist {
			res, err := http.Get("https://www.mercari.com" + next_url)
			if err != nil {
				log.Println(err)
			}
			defer res.Body.Close()
			doc, _ := goquery.NewDocumentFromReader(res.Body)
			next_url, exist = doc.Find(".pager-next .pager-cell a").Attr("href")
			if exist {
				url_list = append(url_list, "https://www.mercari.com"+next_url)
			}
		}
	}

	return url_list
}

func Get_details_item(url string) model.Item_info_mercari {
	res, err := http.Get(url)
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
	if len(no) == 0 {
		no = "NON"
	}

	var soldout bool
	sold := doc.Find(".item-box-container .item-sold-out-badge")
	if sold.Text() == "" {
		soldout = false
	} else {
		soldout = true
	}

	title := doc.Find(".item-name").First().Text()

	price := doc.Find(".item-price-box .item-price").First().Text()

	item_info := model.Item_info_mercari{
		Product_number: no,
		Name:           title,
		Url:            url,
		Sold:           soldout,
		Price:          price,
	}

	return item_info
}

func process_m(url string, userid string, ch chan []model.Item_info_mercari) {
	log.Println("process start", url)
	var list []model.Item_info_mercari
	items_url := Get_items_url(url)
	for _, s := range items_url {
		item_info := Get_details_item(s)
		item_info.Username = userid
		list = append(list, item_info)
	}

	ch <- list

}

func Get_items_on_mercari(userid string, username string) []model.Item_info_mercari {
	var res []model.Item_info_mercari

	url_list := Get_all_page_url(userid)

	ch := make([]chan []model.Item_info_mercari, len(url_list))
	for i, v := range url_list {
		ch[i] = make(chan []model.Item_info_mercari)
		go process_m(v, username, ch[i])
	}

	for i := range ch {
		res = append(res, <-ch[i]...)
		close(ch[i])
	}
	// items_url := Get_items_url(url_list)
	// log.Println("Scraping START", "username:", username, "個数:", strconv.Itoa(len(items_url)))

	// for _, s := range items_url {
	// 	item_info := Get_details_item(s)
	// 	item_info.Username = username
	// 	log.Println(item_info)
	// 	res = append(res, item_info)
	// }

	return res
}
