package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

func Register(username, password string) error {
	client, err := mongo.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("your_database_name").Collection("users")

	// Check if the username already exists
	var existingUser User
	err = collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&existingUser)
	if err == nil {
		return fmt.Errorf("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a new user
	newUser := User{
		Username: username,
		Password: string(hashedPassword),
	}

	_, err = collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Register("testuser", "testpassword")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User registered successfully!")
}
