package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mmuoDev/commons/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo"
)

func main() {
	port := "9000"
	os.Setenv("MONGO_URL", "mongodb://localhost:27017")
	os.Setenv("MONGO_DB_NAME", "commons")
	provideDB, err := mongo.NewConfigFromEnvVars().ToProvider(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Starting server on port:%s", port))
	mongo.NewCollection(provideDB, "cattle")

	// Set client options
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// // Connect to MongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Check the connection
	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to MongoDB!")

	// collection := client.Database("test").Collection("trainers")
}
