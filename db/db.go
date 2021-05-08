package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Item_info struct {
	Product_number string `dynamo:"Product_number"`
	Test           int    `dynamo:"Test_no"`
	Text           string `dynamo:"Text"`
}

type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}

func Create() {
	// dynamoDbRegion := os.Getenv("AWS_REGION")
	dynamoDbRegion := "ap-northeast-1"
	disableSsl := false

	// dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	dynamoDbEndpoint := "http://localhost:8000"
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})

	table := db.Table("mercari_items")

	item := Item_info{
		Product_number: "NT116",
		Test:           1123,
		Text:           "oehfwoeihfwoihe",
	}

	if err := table.Put(item).Run(); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
	}

	var results []Item_info

	if err := table.Scan().All(&results); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
	}

	fmt.Println(results)

}
