package model

type Item_info struct {
	Product_number string `dynamo:"Product_number"`
	Test           int    `dynamo:"Test_no"`
	Text           string `dynamo:"Text"`
}
