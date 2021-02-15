package main

import (
	"log"
	"os"
	// "context"
	// "fmt"
	// "log"
	// "os"
	// "github.com/mmuoDev/commons/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type Partner struct {
	ID      string
	Name    string
	Address string
}

func main() {
	//port := "9000"
	os.Setenv("MONGO_URL", "mongodb://localhost:27017")
	os.Setenv("MONGO_DB_NAME", "commons")
	// provideDB, err := mongo.NewConfigFromEnvVars().ToProvider(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(fmt.Sprintf("Starting server on port:%s", port))

	//ash := Partner{"12345", "Jon Rose", "Texas, US"}
	// misty := Partner{"67890", "nnamdi kanu", "Enugu, Nigeria"}

	// partners := []interface{}{ash, misty}

	//var partner Partner
	// col := mongo.NewCollection(provideDB, "cattle")
	// count, err1 := col.CountDocuments(bson.D{{id: "12345"}})
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }
	// log.Println(count)

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

	myMap := map[string]string{"name": "uche"}
	for key, value := range myMap {
		log.Println(key, value)
	}
}
