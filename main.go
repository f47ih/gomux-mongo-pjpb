package main

import (
    "fmt"
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetClient() *mongo.Client {
    clientOptions := options.Client().ApplyURI("mongodb+srv://dbAdmin:demuji@pjpb.8gm8p.mongodb.net/lapor?retryWrites=true&w=majority")
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Connect(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    return client
}

func main() {
	fmt.Println("Hello")
	c := GetClient()
    err := c.Ping(context.Background(), readpref.Primary())
    if err != nil {
        log.Fatal("Couldn't connect to the database", err)
    } else {
        log.Println("Connected!")
    }
}