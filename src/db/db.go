package db

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TableInfo struct {
	DBClient  *dynamodb.Client
	TableName string
}

// check table existance
func (tableInfo TableInfo) CheckTableExists(tableName string) (*dynamodb.DescribeTableOutput, error) {
	result, err := tableInfo.DBClient.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{TableName: &tableName})

	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", tableName)
			err = nil
		}
	}
	return result, nil
}

// get all beds data
func (tableInfo TableInfo) GetAllBeds() ([]map[string]types.AttributeValue, error) {
	result, err := tableInfo.DBClient.Scan(context.TODO(), &dynamodb.ScanInput{TableName: &tableInfo.TableName})

	if err != nil {
		log.Fatalln("fatal error", err)
		return nil, err
	}
	return result.Items, nil
}

// get bed availability of a particular type of category i.e. general/ICU/CCU
func (tableInfo TableInfo) GetBedDetails(bedTypeId string) (map[string]types.AttributeValue, error) {
	projectionList := "bed_type, t_capacity, occupied, available"
	result, err := tableInfo.DBClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &tableInfo.TableName,
		Key: map[string]types.AttributeValue{
			"bed_type_id": &types.AttributeValueMemberS{Value: bedTypeId},
		},
		ProjectionExpression: &projectionList,
	})

	if err != nil {
		log.Fatalln("fatal error", err)
		return nil, err
	}
	return result.Item, nil
}
