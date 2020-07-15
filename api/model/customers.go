package model

import (
	"context"
	"encoding/json"	
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CUSTOMER ...
type CUSTOMER struct {
	UserID      string `json:"user_id"`
	Login       string `json:"login"`
	Password    int    `json:"password"`
	Name        string `json:"name"`
	CompanyID   int    `json:"company_id"`
	CreditCards string `json:"credit_cards"`
}


// CUSTOMERResponse ...
type CUSTOMERResponse struct {
	CUSTOMER    CUSTOMER               `json:"customer"`
}

// CUSTOMERsResponse ...
type CUSTOMERsResponse struct {
	CUSTOMERs []CUSTOMERResponse `json:"customers"`
}

// CUSTOMERsGetAll ...
func CUSTOMERsGetAll(coll *mongo.Collection, findLimit int64) (CUSTOMERsResponse, error) {
	i := CUSTOMERsResponse{CUSTOMERs: []CUSTOMERResponse{}}
	findOptions := options.Find()
	findOptions.SetLimit(findLimit)
	results := []CUSTOMERResponse{}
	log.Println("Executing find for CUSTOMERs...")
	cur, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Iterating using cursor...")
	for cur.Next(context.TODO()) {
		var elem CUSTOMERResponse
		err := cur.Decode(&elem)
		if err != nil {
			log.Println("Error decoding with cursor, err:", err)
			return i, err
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Println("Error using find cursor, err:", err)
		return i, err
	}
	cur.Close(context.TODO())

	i.CUSTOMERs = results

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return i, nil
}

// CUSTOMERDeleteByID ...
func CUSTOMERDeleteByID(coll *mongo.Collection, id string) error {
	log.Println("Deleting CUSTOMER id:", id)
	idPrim, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	r, err := coll.DeleteOne(context.TODO(), bson.M{"_id": idPrim})
	if err != nil {
		return err
	}
	count := r.DeletedCount
	log.Println("Deleted count:", count)
	return nil
}

// CUSTOMERDeleteAll ...
func CUSTOMERDeleteAll(coll *mongo.Collection) error {
	log.Println("Deleting all CUSTOMERs...")
	r, err := coll.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	count := r.DeletedCount
	log.Println("Deleted count:", count)
	return err
}

// CUSTOMERGetByID ...
func CUSTOMERGetByID(coll *mongo.Collection, id string) (CUSTOMERResponse, error) {
	i := CUSTOMERResponse{}
	idPrim, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return i, err
	}
	coll.FindOne(context.TODO(), bson.M{"_id": idPrim}).Decode(&i)
	return i, nil
}

// CUSTOMERInsert ...
func CUSTOMERInsert(
	coll *mongo.Collection,
	customer CUSTOMER,
) (CUSTOMERResponse, error) {
	var response CUSTOMERResponse
	customerBytes, err := json.Marshal(customer)
	if err != nil {
		log.Println("Error marshalling customer to bytes.")
		return response, err
	}
	json.Unmarshal(customerBytes, &customer)

	customerResponse := CUSTOMERResponse{CUSTOMER: customer}
	r, err := coll.InsertOne(context.Background(), customerResponse)
	if err != nil {
		log.Println("Error inserting CUSTOMER:", err)
		return response, err
	}
	log.Println("Inserted CUSTOMER with id:", r.InsertedID.(primitive.ObjectID).Hex())
	return response, err
}

