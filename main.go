package main

import (
	"log"
	"os"

	"github.com/mmuoDev/commons/mysql"
	// "context"
	// "fmt"
	// "log"
	// "os"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type Partner struct {
	ID      string
	Name    string
	Address string
}

func main() {
	//port := "9000"
	os.Setenv("MYSQL_USERNAME", "root")
	os.Setenv("MYSQL_PASSWORD", "@Password12")
	os.Setenv("MYSQL_HOST", "127.0.0.1:3306")
	os.Setenv("MYSQL_DB_NAME", "pangaea")
	provideDB, err := mysql.NewConfigFromEnvVars().ToProvider()
	if err != nil {
		log.Fatal(err)
	}

	// query := "CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text, product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)"

	// err1 := provideDB.Create(query)
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }
	// log.Println("created!")

	student := map[string]interface{}{"product_name": "door", "product_price": 5000}

	count, err := provideDB.Insert("product", student)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(count)
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
