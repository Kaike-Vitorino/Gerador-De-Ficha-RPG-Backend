package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// UploadCharacterSheet faz upload da ficha de personagem e suas informações no MongoDB
func UploadCharacterSheet(client *mongo.Client, dbName, collectionName string, data interface{}) error {
	collection := client.Database(dbName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	log.Println("Character sheet uploaded successfully!")
	return nil
}
