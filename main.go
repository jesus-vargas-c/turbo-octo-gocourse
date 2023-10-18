package main

import (
	"fmt"
)

// main is the entry point for the program
func main() {
	// //getting config from aws_manager.go
	// cfg, err := ReturnLoadConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //getting dynamodb client based on config
	// svc := ReturnDynamoClient(cfg)

	// //list tables in dynamodb
	// tables, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
	// 	Limit: aws.Int32(5),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(tables)
	for i := 0; i < 10000; i++ {
		fmt.Println(i + 1)
	}
}
