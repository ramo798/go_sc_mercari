package model

type Item_info_mercari struct {
	Product_number string `dynamo:"Product_number,hash"  csv:"品番" dynamodbav:"Product_number"`
	Name           string `dynamo:"Name" csv:"物品名" dynamodbav:"Name"`
	Url            string `dynamo:"Url" csv:"Url" dynamodbav:"Url"`
	Sold           bool   `dynamo:"Sold" csv:"SOLD" dynamodbav:"Sold"`
	Username       string `dynamo:"Username" csv:"Username" dynamodbav:"Username"`
	Duplicate      bool   `dynamo:"Duplicate" csv:"Duplicate" dynamodbav:"Duplicate"`
}
