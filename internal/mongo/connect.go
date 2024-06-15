package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Client *mongo.Client

// ConnectMongoDB inicializa a conexão com o MongoDB
func ConnectMongoDB(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	Client = client
	log.Println("Connected to MongoDB!")
	return Client, nil
}

// GetCollection retorna uma coleção específica do MongoDB
func GetCollection(database, collection string) *mongo.Collection {
	return Client.Database(database).Collection(collection)
}
