package main

import (
	"log"
	"strconv"
	"test/crawler"
	"test/csvmaker"
	"test/db"
)

func main() {
	mercari_items := crawler.Get_items_on_mercari("951762445", "maron")
	log.Println("maron items: ", len(mercari_items))
	log.Println("Writning START", "maron")
	for index, i := range mercari_items {
		if !i.Sold {
			if i.Product_number == "NON" {
				i.Product_number = "NN" + strconv.Itoa(index)
			}
			log.Println("writing", i)
			db.Create(i, "mercari_items")
		}

	}
	db.Scan("mercari_items")

	log.Println("Writning START", "younghoho_1121")
	yi := crawler.Get_items_on_yahuoku("younghoho_1121")
	log.Println("younghoho_1121 items: ", len(yi))
	for index, i := range yi {
		if !i.Sold {
			if i.Product_number == "NON" {
				i.Product_number = "NN" + strconv.Itoa(index)
			}
			// log.Println("writing", i)
			db.Create(i, "yahuoku_items")
		}
	}

	log.Println("Writning START", "tomokimi_777")
	yj := crawler.Get_items_on_yahuoku("tomokimi_777")
	log.Println("tomokimi_777 items: ", len(yj))
	for index, i := range yj {
		if !i.Sold {
			if i.Product_number == "NON" {
				i.Product_number = "NN" + strconv.Itoa(index)
			}
			// log.Println("writing", i)
			db.Create(i, "yahuoku_items")
		}
	}

	// // db.Scan("yahuoku_items")

	csvmaker.Makecsv()

	// crawler.Get_items_on_yahuoku("tomokimi_777")
}
