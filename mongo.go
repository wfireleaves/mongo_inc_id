package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URL        = "mongodb://172.24.128.54:27017"
	DATABASE   = "counter"
	COLLECTION = "counter"
)

func GetMongoCollection() (*mongo.Collection, error) {
	opts := options.Client()
	opts.ApplyURI(URL)
	opts.SetConnectTimeout(time.Second * 3)
	opts.SetMaxPoolSize(5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("connect to mongo failed,err:%s", err.Error())
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping mongo failed,err:%s", err.Error())
	}
	collection := client.Database(DATABASE).Collection(COLLECTION)
	if collection == nil {
		return nil, fmt.Errorf("collection not exist,database:%s collection:%s", DATABASE, COLLECTION)
	}
	return collection, nil
}
