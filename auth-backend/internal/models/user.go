package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const db string = "WebChatApp"
const collection string = "Users"

type User struct {
	Username string	
	HashedPassword string 
}

type UserModel struct {
	Client *mongo.Client
}

func (m *UserModel) Insert(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	exists, err := m.Exists(username)
	if err != nil {
		return err
	}

	if exists {
		return ErrDuplicateUsername
	}

	coll := m.Client.Database(db).Collection(collection)
	doc := User{Username: username, HashedPassword: string(hashedPassword)}

	_, err = coll.InsertOne(context.TODO(), doc)

	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Exists(username string) (bool, error) {
	coll := m.Client.Database(db).Collection(collection)
	filter := bson.M{"username": username}
	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	return count >= 1, nil
}

type UserMongoFields struct {
	Username string `bson:"username"`
	HashedPassword string `bson:"hashedpassword"`
	ID primitive.ObjectID `bson:"_id"`
}

func (m *UserModel) Authenticate(username, password string) (string, error) {
	coll := m.Client.Database(db).Collection(collection)
	filter := bson.M{"username": username}
	var user UserMongoFields
	result := coll.FindOne(context.TODO(), filter)
	err := result.Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", ErrNoRecord
		}
		return "" , err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword){
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	return user.ID.Hex(), nil

}