package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

//main is the entry point for the program
func main() {
	//getting config from aws_manager.go
	cfg, err := ReturnLoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	//getting dynamodb client based on config
	svc := ReturnDynamoClient(cfg)

	//list tables in dynamodb
	tables, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tables)

}
