package main

import (
	"fmt"
	"test/crawler"
	"time"
)

func main() {
	start := time.Now()
	mercari_items := crawler.Get_items_on_mercari("951762445")
	fmt.Println(mercari_items)
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
