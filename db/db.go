package db

import (
	"fmt"
	"test/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var db *dynamo.DB

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		// Endpoint:    aws.String(os.Getenv("DYNAMO_ENDPOINT")),
		Endpoint:    aws.String("http://dynamodb-local:8000"),
		Credentials: credentials.NewStaticCredentials("fakeMyKeyId", "fakeSecretAccessKey", ""),
	})
	if err != nil {
		panic(err)
	}
	db = dynamo.New(sess)
}

// Create() はmodel.Item_infoを渡すとPUTする関数
func Create(item model.Item_info_mercari, table_name string) {
	table := db.Table(table_name)
	err := table.Put(item).Run()
	if err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
	}
}

func Scan(table_name string) []model.Item_info_mercari {
	table := db.Table(table_name)
	var results []model.Item_info_mercari
	if err := table.Scan().All(&results); err != nil {
		fmt.Printf("Failed to scan item[%v]\n", err)
	}
	return results
}
