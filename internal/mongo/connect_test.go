package mongo

import (
	"testing"
)

func TestConnectMongoDB(t *testing.T) {
	uri := "your-mongo-uri"
	client, err := ConnectMongoDB(uri)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	if client == nil {
		t.Fatal("Expected client to be non-nil")
	}
}
