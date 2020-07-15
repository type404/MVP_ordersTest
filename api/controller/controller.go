package controller

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Controller ...
type Controller struct {
	MongoClient              *mongo.Client
	MongoDBCollCompany       *mongo.Collection
	MongoDBCollCustomer      *mongo.Collection
	MongoFindLimit      	 int64
}

// NewController example
func NewController(
	mongoDBConnStr string,
	mongoDBName string,
	mongoDBCollCompany string,
	mongoDBCollCustomer string,
	mongoFindLimit int64,
) (*Controller, error) {
	mc, err := createMongoDbConnection(mongoDBConnStr)
	if err != nil {
		return nil, err
	}
	collcompany := createMDbCollection(mc, mongoDBName, mongoDBCollCompany)
	collcustomer := createMDbCollection(mc, mongoDBName, mongoDBCollCustomer)

	c := Controller{
		MongoClient:         mc,
		MongoDBCollCompany:  collcompany,
		MongoDBCollCustomer: collcustomer,
		MongoFindLimit:      mongoFindLimit,
	}
	return &c, nil
}

func createMDbCollection(
	client *mongo.Client,
	databaseName string,
	collectionName string,
) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

func createMongoDbConnection(mongoDBConnStr string) (*mongo.Client, error) {
	log.Println("Creating MongoDB Connection...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoDBConnStr).SetDirect(true)
	c, err := mongo.NewClient(clientOptions)

	err = c.Connect(ctx)

	if err != nil {
		log.Println("Unable to initialize connection with error:", err)
		return nil, err
	}

	log.Println("Trying to ping MongoDB...")
	err = c.Ping(ctx, nil)
	if err != nil {
		log.Println("Unable to ping MongoDB with error:", err)
		return nil, err
	}
	return c, nil
}
