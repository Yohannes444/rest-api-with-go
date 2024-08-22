package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clinetInstaceerror error
type Collection string
const( 
	ProductCollection Collection = "products"
	UserCollection Collection = "users"
)
var mongoOnce sync.Once
const (
	uri= "mongodb://localhost:27017"
	Database = "products-api"
)
func GetMongoClient()(*mongo.Client, error){
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.TODO(),clientOptions)

		 clientInstance = client
		 clinetInstaceerror = err
	})

	return clientInstance, clinetInstaceerror
}