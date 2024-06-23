package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Login(username, password string) (bool, error) {
	client, err := mongo.Connect()
	if err != nil {
		return false, err
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("your_database_name").Collection("users")

	// Find the user
	var user User
	err = collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return false, fmt.Errorf("username or password incorrect")
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, fmt.Errorf("username or password incorrect")
	}

	return true, nil
}

func main() {
	success, err := Login("testuser", "testpassword")
	if err != nil {
		log.Fatal(err)
	}
	if success {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed!")
	}
}
