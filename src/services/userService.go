package services

import (
	"context"
	"fmt"

	"github.com/hamzapro305/GoJwt/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(email string, password string) error {
	// encrupt pass
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	// Add user to database
	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: hashedPassword,
	}
	_, err = models.UserCollection.InsertOne(context.TODO(), newUser)
	return err
}

func GetUser(email string) (models.User, error) {
	filter := bson.M{"email": email}
	var user models.User
	err := models.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}

func ValidateUser(email string, password string) (*models.User, error) {
	// Fetch user by email
	user, err := GetUser(email)
	if err != nil {
		return nil, fmt.Errorf("user not found") // Return an error, not an HTTP response
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password incorrect") // Return an error, not an HTTP response
	}

	// Return the user if everything is fine
	return &user, nil
}
