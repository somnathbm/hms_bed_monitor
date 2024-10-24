package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sso/types"
)

type TableBasics struct {
	DyanamoDBClient *dynamodb.Client
	TableName       string
}

func (tableBasics TableBasics) CheckTableExists(ctx context.Context) (bool, error) {
	exists := true
	_, err := tableBasics.DyanamoDBClient.DescribeTable(ctx, &dynamodb.DescribeTableInput{TableName: aws.String(tableBasics.TableName)})

	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("The table %v does not exist", tableBasics.TableName)
			err = nil
		} else {
			log.Printf("Could not determine the root cause. Here is more info: %v", err)
		}
		exists = false
	}
	return exists, err
}

func main() {
	basics := TableBasics{
		TableName:       "hospi_bed_test_table",
		DyanamoDBClient: dynamodb.New(dynamodb.Options{Region: "us-east-1"}),
	}

	_, err := basics.CheckTableExists(context.TODO())

	if err != nil {
		fmt.Printf("Operation not succeed due to %v", err)
	}
	fmt.Print("The table does exists!")
}
