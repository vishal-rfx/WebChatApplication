package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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