package csvmaker

import (
	"os"
	"test/db"

	"github.com/gocarina/gocsv"
)

func Makecsv() {
	yahuoku := db.Scan("yahuoku_items")
	mercari := db.Scan("mercari_items")

	file, _ := os.OpenFile("yahuoku.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()

	file2, _ := os.OpenFile("mercari.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()

	gocsv.MarshalFile(&yahuoku, file)
	gocsv.MarshalFile(&mercari, file2)
}

func Jointcsv() {

}
