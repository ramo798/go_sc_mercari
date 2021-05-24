package main

import (
	"fmt"
	"log"
	"strconv"
	"test/crawler"
	"test/csvmaker"
	"test/db"
	"test/model"
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
	db.Create_table("mercari_items")
	db.Create_table("yahuoku_items")

	log.Println("START", "maron")
	mercari_items := crawler.Get_items_on_mercari("951762445", "maron")
	log.Println("maron items: ", len(mercari_items))

	uniq, n_unic := unique(mercari_items)
	for index, i := range uniq {
		if !i.Sold {
			if i.Product_number == "NON" {
				i.Product_number = "NN" + strconv.Itoa(index)
			}
			log.Println("writing", i)
			db.Create(i, "mercari_items")
		}
	}
	for index, j := range n_unic {
		if !j.Sold {
			if j.Product_number == "NON" {
				j.Product_number = "NN" + strconv.Itoa(index)
			}
			j.Product_number = strconv.Itoa(index) + "+" + j.Product_number
			log.Println("writing", j)
			db.Create(j, "mercari_items")
		}
	}

	var y_items []model.Item_info_mercari

	log.Println("START", "younghoho_1121")
	yi := crawler.Get_items_on_yahuoku("younghoho_1121")
	y_items = append(y_items, yi...)
	log.Println("START", "tomokimi_777")
	yj := crawler.Get_items_on_yahuoku("tomokimi_777")
	y_items = append(y_items, yj...)

	uniq, n_unic = unique(y_items)

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

type MyEvent struct {
	Name string `json:"What is your name?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func hello(event MyEvent) (MyResponse, error) {
	division()
	return MyResponse{Message: fmt.Sprintf("Hello %s!!", event.Name)}, nil
}

func main() {
	division()
	csvmaker.Makecsv()

	// lambda.Start(hello)

}
