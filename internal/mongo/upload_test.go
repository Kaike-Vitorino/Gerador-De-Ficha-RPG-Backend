package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestUploadCharacterSheet(t *testing.T) {
	uri := "your-mongo-uri"
	client, err := ConnectMongoDB(uri)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	data := bson.D{{"name", "Test Character"}, {"level", 1}}
	err = UploadCharacterSheet(client, "your-database", "your-collection", data)
	if err != nil {
		t.Fatalf("Failed to upload character sheet: %v", err)
	}
}
