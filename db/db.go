package db

import (
	"fmt"
	"os"
	"test/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// Create() はmodel.Item_infoを渡すとPUTする関数
func Create(item model.Item_info_mercari) {
	dynamoDbRegion := os.Getenv("AWS_REGION")
	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	disableSsl := false

	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	} else {
		dynamoDbEndpoint = "http://localhost:8000"

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

	if err := table.Put(item).Run(); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
	}

}

func Scan() {
	dynamoDbRegion := os.Getenv("AWS_REGION")
	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	disableSsl := false

	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	} else {
		dynamoDbEndpoint = "http://localhost:8000"

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
	var results []model.Item_info_mercari

	if err := table.Scan().All(&results); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
	}

	fmt.Println(results)
}
