package main

import (
	"log"
	"strconv"
	"test/crawler"
	"test/db"
	"test/model"
	"time"
)

func unique(list []model.Item_info_mercari) ([]model.Item_info_mercari, []model.Item_info_mercari) {
	var res_u []model.Item_info_mercari
	var res_n []model.Item_info_mercari
	m := map[string]int{}
	for _, s := range list {
		m[s.Product_number]++ // 出現回数をカウント
	}
	for i, j := range m {
		if j == 1 {
			for _, k := range list {
				if i == k.Product_number {
					res_u = append(res_u, k)
				}
			}
		} else {
			for _, k := range list {
				if i == k.Product_number {
					res_n = append(res_n, k)
				}
			}
		}
	}
	return res_u, res_n
}

func division() {
	log.Println("START", "maron")
	mercari_items := crawler.Get_items_on_mercari("951762445", "maron")
	log.Println("maron items: ", len(mercari_items))
	for index, i := range mercari_items {
		if !i.Sold {
			if i.Product_number == "NON" {
				i.Product_number = "NN" + strconv.Itoa(index)
			}
			log.Println("writing", i)
			db.Create(i, "mercari_items")
		}

	}
	// db.Scan("mercari_items")

	var y_items []model.Item_info_mercari

	log.Println("START", "younghoho_1121")
	yi := crawler.Get_items_on_yahuoku("younghoho_1121")
	y_items = append(y_items, yi...)
	log.Println("START", "tomokimi_777")
	yj := crawler.Get_items_on_yahuoku("tomokimi_777")
	y_items = append(y_items, yj...)

	uniq, n_unic := unique(y_items)

	for index, j := range uniq {
		if !j.Sold {
			if j.Product_number == "NON" {
				j.Product_number = "NN" + strconv.Itoa(index)
			}
			log.Println("writing", j)
			db.Create(j, "yahuoku_items")
		}
	}

	for index, j := range n_unic {
		if !j.Sold {
			if j.Product_number == "NON" {
				j.Product_number = "NN" + strconv.Itoa(index)
			}
			j.Product_number = strconv.Itoa(index) + "+" + j.Product_number
			log.Println("writing", j)
			db.Create(j, "yahuoku_items")
		}
	}

	// log.Println("younghoho_1121 items: ", len(yi))
	// for index, i := range yi {
	// 	if !i.Sold {
	// 		if i.Product_number == "NON" {
	// 			i.Product_number = "NN" + strconv.Itoa(index)
	// 		}
	// 		log.Println("writing", i)
	// 		db.Create(i, "yahuoku_items")
	// 	}
	// }

	// log.Println("START", "tomokimi_777")
	// yj := crawler.Get_items_on_yahuoku("tomokimi_777")
	// log.Println("tomokimi_777 items: ", len(yj))
	// for index, i := range yj {
	// 	if !i.Sold {
	// 		if i.Product_number == "NON" {
	// 			i.Product_number = "NN" + strconv.Itoa(index)
	// 		}
	// 		log.Println("writing", i)
	// 		db.Create(i, "yahuoku_items")
	// 	}
	// }
}

func wtest() {
	var test []model.Item_info_mercari
	i := model.Item_info_mercari{
		Product_number: "aaa",
		Username:       "testman",
	}
	j := model.Item_info_mercari{
		Product_number: "bbb",
		Username:       "tesascascascascastman",
	}
	k := model.Item_info_mercari{
		Product_number: "aaa",
		Username:       "testsacascascman",
	}
	// db.Create_s(i, "yahuoku_items")
	// fmt.Println(db.Scan("yahuoku_items"))
	// fmt.Println(i.Product_number)
	test = append(test, i)
	test = append(test, j)
	test = append(test, k)
	unique(test)

}

func main() {
	time.Sleep(time.Second * 10)
	log.Println("start")
	// wtest()
	division()
	// fmt.Println(db.Scan("yahuoku_items"))
	// fmt.Println(len(db.Scan("yahuoku_items")))
	log.Println("stop")

	// csvmaker.Makecsv()

}
