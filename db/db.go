package db

import (
	"fmt"
	"log"
	"os"
	"test/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/guregu/dynamo"
)

var db *dynamo.DB
var db2 *dynamodb.DynamoDB

func init() {
	var sess *session.Session
	var err error

	if len(os.Getenv("DYNAMO_ENDPOINT")) != 0 {
		sess, err = session.NewSession(&aws.Config{
			Region:      aws.String("ap-northeast-1"),
			Endpoint:    aws.String(os.Getenv("DYNAMO_ENDPOINT")),
			Credentials: credentials.NewStaticCredentials("fakeMyKeyId", "fakeSecretAccessKey", ""),
		})
		if err != nil {
			panic(err)
		}
	} else {
		sess, err = session.NewSession(&aws.Config{
			Region:      aws.String("ap-northeast-1"),
			Credentials: credentials.NewStaticCredentials(os.Getenv("DYNAMO_ACCESS"), os.Getenv("DYNAMO_SECLET"), ""),
		})
		if err != nil {
			panic(err)
		}
	}

	db = dynamo.New(sess)
	db2 = dynamodb.New(sess)

}

func Create_table(table_name string) {
	db.Table(table_name).DeleteTable().Run()
	err := db.CreateTable(table_name, model.Item_info_mercari{}).Run()
	if err != nil {
		panic(err)
	}
}

// Create() はmodel.Item_infoを渡すとPUTする関数
func Create(item model.Item_info_mercari, table_name string) {
	table := db.Table(table_name)
	err := table.Put(item).Run()
	if err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
	}
}

func Create_s(item model.Item_info_mercari, table_name string) {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(table_name),
		Item:      av,
	}
	_, err = db2.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	log.Println(11111)
}

func Scan(table_name string) []model.Item_info_mercari {
	table := db.Table(table_name)
	var results []model.Item_info_mercari
	if err := table.Scan().All(&results); err != nil {
		fmt.Printf("Failed to scan item[%v]\n", err)
	}
	return results
}
