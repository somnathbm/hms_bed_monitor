package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	// "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"hospi_bed_stats/models"
)

// Get mongo collection for 1 minute connection pool
func get_db_collection() (*mongo.Collection, *mongo.Client) {
	_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("APP_DB_URI")))
	if err != nil {
		panic(err)
	}
	collection := client.Database(os.Getenv("APP_DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	return collection, client
}

// Get all beds
func GetAllBeds() ([]models.Beds, error) {
	var results []models.Beds
	var decodedResult []models.Beds

	collection, client := get_db_collection()
	if collection == nil || client == nil {
		// panic("unable to proceed operation with mongo")
		return nil, errors.New("could not get db")
	}

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil && err == mongo.ErrNoDocuments {
		// handle error
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("ERROR Ocurred!!")
		panic(err)
	}
	// Prints the results of the find operation as structs
	for _, result := range results {
		cursor.Decode(&result)
		decodedResult = append(decodedResult, result)
	}
	client.Disconnect(context.Background())
	return decodedResult, nil
}

// get bed availability of a particular type of category i.e. general/ICU/CCU
func GetBedDetails(bedTypeId string) (primitive.M, error) {
	collection, client := get_db_collection()
	if collection == nil || client == nil {
		// panic("unable to proceed operation with mongo")
		return nil, errors.New("could not get db")
	}
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"bed_type_id": bedTypeId}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&result)
	if err != nil && err == mongo.ErrNoDocuments {
		// handle error
		return nil, err
	}
	client.Disconnect(context.Background())
	return result, nil
}
