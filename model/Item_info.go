package model

type Item_info_mercari struct {
	Product_number string `dynamo:"Product_number"`
	Name           string `dynamo:"Name"`
	Url            string `dynamo:"Url"`
	Sold           bool   `dynamo:"Sold"`
}
