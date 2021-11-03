package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var MongoConnection = ConnectionDB()
var clientoptions = options.Client().ApplyURI("mongodb+srv://root:<")

func ConnectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientoptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion Exitosa con la BD")

	return client
}

func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}

func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
