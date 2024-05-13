package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStore struct {
	coll *mongo.Collection
}

func NewUserStore(store *MongoStore) *UserStore {
	coll := store.GetCollection("users")
	index := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := coll.Indexes().CreateOne(context.Background(), index, nil)
	if err != nil {
		log.Println(err)
	}
	return &UserStore{coll: coll}
}

func (us *UserStore) SaveUser(user *UserDao) (*UserDao, error) {
	result, err := us.coll.InsertOne(context.Background(), user.BsonD())
	if err != nil {
		return nil, err
	}
	var insertedId primitive.ObjectID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		insertedId = oid
	}
	user.ID = insertedId
	return user, nil
}

func (us *UserStore) GetUserByEmail(email string) (*UserDao, error) {
	filter := bson.D{{"email", email}}
	var result *UserDao
	err := us.coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type UserDao struct {
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	ID       primitive.ObjectID `bson:"_id,omitempty"`
}

func (ud *UserDao) BsonD() bson.D {
	return bson.D{{"email", ud.Email}, {"password", ud.Password}, {"_id", ud.ID}}
}
