package main

import (
	"log"
	"strconv"
	"test/crawler"
	"test/db"
	"time"
)

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
			db.Create_s(i, "mercari_items")
		}

	}
	// db.Scan("mercari_items")

	log.Println("START", "younghoho_1121")
	yi := crawler.Get_items_on_yahuoku("younghoho_1121")
	log.Println("younghoho_1121 items: ", len(yi))
	for index, i := range yi {
		if !i.Sold {
			if i.Product_number == "NON" {
				i.Product_number = "NN" + strconv.Itoa(index)
			}
			log.Println("writing", i)
			db.Create_s(i, "yahuoku_items")
		}
	}

	log.Println("START", "tomokimi_777")
	yj := crawler.Get_items_on_yahuoku("tomokimi_777")
	log.Println("tomokimi_777 items: ", len(yj))
	for index, i := range yj {
		if !i.Sold {
			if i.Product_number == "NON" {
				i.Product_number = "NN" + strconv.Itoa(index)
			}
			log.Println("writing", i)
			db.Create_s(i, "yahuoku_items")
		}
	}
}

func main() {
	time.Sleep(time.Second * 10)
	log.Println("start")
	division()

	log.Println("stop")

	// csvmaker.Makecsv()

}
