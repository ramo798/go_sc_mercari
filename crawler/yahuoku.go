package crawler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"test/model"

	"github.com/PuerkitoBio/goquery"
)

// duplicate_deletion is リストの重複を削除する
func duplicate_deletion(list []string) []string {
	m := make(map[string]struct{})

	newList := make([]string, 0)

	for _, element := range list {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; !ok {
			m[element] = struct{}{}
			newList = append(newList, element)
		}
	}
	return newList
}

func Get_items_url_y(urls []string) []string {
	var resUrl []string

	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		defer res.Body.Close()

		doc, _ := goquery.NewDocumentFromReader(res.Body)
		doc.Find(".inner a").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			if url != "#dummyWatch" && url != "https://auctions.yahoo.co.jp/search/advanced?" {
				resUrl = append(resUrl, url)
				// fmt.Println(url)
			}
		})

	}
	resUrl = duplicate_deletion(resUrl)

	return resUrl
}

func Get_all_page_url_y(uid string) []string {
	origin_url := "https://auctions.yahoo.co.jp/seller/" + uid

	url_list := []string{origin_url}

	res, err := http.Get(origin_url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	next_url, exist := doc.Find("#ASsp1 > p.next > a").Attr("href")

	if exist {
		url_list = append(url_list, next_url)
		for exist {
			res, err := http.Get(next_url)
			if err != nil {
				log.Println(err)
			}
			defer res.Body.Close()
			doc, _ := goquery.NewDocumentFromReader(res.Body)
			next_url, exist = doc.Find("#ASsp1 > p.next > a").Attr("href")
			if exist {
				url_list = append(url_list, next_url)
			}
		}
	}

	return url_list
}

func Get_details_item_y(url string) model.Item_info_mercari {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	text := doc.Find("#adoc > div.ProductExplanation__body.highlightWordSearch > div.ProductExplanation__commentArea > div > center > table > tbody > tr:nth-child(6) > td:nth-child(2) > font").First().Text()

	var no string
	if strings.Contains(text, "[") && strings.Contains(text, "]") {
		no = text[strings.Index(text, "[")+1 : strings.Index(text, "]")]
	}
	if len(no) == 0 {
		no = "NON"
	}
	// log.Println(no)

	var soldout bool = false

	title := doc.Find(".ProductTitle__title .ProductTitle__text").First().Text()
	// log.Println("title", title)

	item_info := model.Item_info_mercari{
		Product_number: no,
		Name:           title,
		Url:            url,
		Sold:           soldout,
	}

	return item_info
}

func Get_items_url_y_for_process(url string) []string {
	var resUrl []string

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".inner a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		if url != "#dummyWatch" && url != "https://auctions.yahoo.co.jp/search/advanced?" {
			resUrl = append(resUrl, url)
		}
	})

	resUrl = duplicate_deletion(resUrl)

	return resUrl
}

func process(url string, userid string, ch chan []model.Item_info_mercari) {
	// defer wg.Done()
	log.Println(2)
	var list []model.Item_info_mercari
	items_url := Get_items_url_y_for_process(url)
	for _, s := range items_url {
		item_info := Get_details_item_y(s)
		// log.Println(item_info)
		item_info.Username = userid
		list = append(list, item_info)
	}
	log.Println("in process")
	ch <- list

}

func Get_items_on_yahuoku(userid string) []model.Item_info_mercari {
	var res []model.Item_info_mercari

	url_list := Get_all_page_url_y(userid)

	// wg := new(sync.WaitGroup)
	ch := make([]chan []model.Item_info_mercari, len(url_list))
	for i, v := range url_list {
		// wg.Add(1)
		ch[i] = make(chan []model.Item_info_mercari)
		go process(v, userid, ch[i])
	}
	// wg.Wait()
	fmt.Println(2222222, ch)
	for i, _ := range ch {
		// fmt.Println(<-ch[i])
		// for _, s := range <-ch[i] {
		// res = append(res, s)
		// }
		res = append(res, <-ch[i]...)
		close(ch[i])
	}

	// close(ch)
	fmt.Println("Finish!")

	fmt.Println(res)
	fmt.Println(len(res))
	return res
}
