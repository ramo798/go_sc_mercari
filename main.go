package main

import (
	"fmt"
	"test/crawler"
)

func main() {
	mercari_items := crawler.Get_items_on_mercari("951762445")
	fmt.Println(mercari_items)

}
