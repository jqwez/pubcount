package store

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	Client       *mongo.Client
	databaseName string
}

func NewMongoStore() *MongoStore {
	ms := &MongoStore{}
	client := ms.MustConnectFromEnv()
	ms.Client = client
	return ms
}

func (ms *MongoStore) SetDatabase(dbName string) *MongoStore {
	ms.databaseName = dbName
	return ms
}

func (ms *MongoStore) Database() string {
	return ms.databaseName
}

func (ms *MongoStore) GetCollection(coll string) *mongo.Collection {
	return ms.Client.Database(ms.databaseName).Collection(coll)
}

func (ms *MongoStore) MustConnectFromEnv() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set MONGODB_URI in environment")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (ms *MongoStore) MustDisconnect() {
	if err := ms.Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}
