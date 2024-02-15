package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type URL struct {
	ID        string    `bson:"_id"`
	Code      string    `bson:"code"`
	URL       string    `bson:"url"`
	CreatedAt time.Time `bson:"created_at"`
}

type MongoDBRepository struct {
	client *mongo.Client
	db     *mongo.Database
	coll   *mongo.Collection
}

func NewMongoDBRepository(uri string, dbName string, collName string) (*MongoDBRepository, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri).SetMaxPoolSize(10))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	coll := db.Collection(collName)

	return &MongoDBRepository{
		client: client,
		db:     db,
		coll:   coll,
	}, nil
}

func (r *MongoDBRepository) Get(code string) (string, error) {
	filter := bson.M{"code": code}
	var result URL

	err := r.coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		return "", err
	}

	return result.URL, nil
}

func (r *MongoDBRepository) Set(code string, url string) error {
	log.Println("set_url: ", url)
	_, err := r.coll.InsertOne(context.Background(), bson.M{
		"code": code,
		"url":  url,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoDBRepository) Close() {
	if r.client != nil {
		r.client.Disconnect(context.Background())
	}
}
