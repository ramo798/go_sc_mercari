package model

type Item_info_mercari struct {
	Product_number string `dynamo:"Product_number"  csv:"品番"`
	Name           string `dynamo:"Name" csv:"物品名"`
	Url            string `dynamo:"Url" csv:"Url"`
	Sold           bool   `dynamo:"Sold" csv:"SOLD"`
	Username       string `dynamo:"Username"csv:"Username"`
}
