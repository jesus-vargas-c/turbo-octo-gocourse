package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Using the SDK's default configuration, loading additional config
// and credentials values from the environment variables, shared
// credentials, and shared configuration `files

func ReturnLoadConfig() (aws.Config, error) {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return cfg, err
}

//return dynamodb client based on config
func ReturnDynamoClient(cfg aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(cfg)

}
