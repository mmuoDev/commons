package main

import (

	// "fmt"
	"log"
	"os"

	"github.com/mmuoDev/commons/mysql"
	// "context"
	// "fmt"
	// "log"
	// "os"
)

type Partner struct {
	ID              string `bson:"id"`
	Name            string `bson:"name"`
	Address         string `bson:"address"`
	MultitenancyKey string `bson:"multitenancykey"`
}

func main() {
	//port := "9000"
	//mongo
	// os.Setenv("MONGO_URL", "mongodb://localhost:27017")
	// os.Setenv("MONGO_DB_NAME", "esusu")
	// provide, err := mongo.NewConfigFromEnvVars().ToProvider(context.Background())
	// if err != nil {
	// 	log.Fatal("unable to connect mongo")
	// }
	// log.Println("connected to DB")
	// col := mongo.NewCollection(provide, "partners")

	os.Setenv("MYSQL_USERNAME", "root")
	os.Setenv("MYSQL_PASSWORD", "@Password12")
	os.Setenv("MYSQL_HOST", "127.0.0.1:3306")
	os.Setenv("MYSQL_DB_NAME", "pangaea")
	provideDB, err := mysql.NewConfigFromEnvVars().ToProvider()
	if err != nil {
		log.Fatal(err)
	}

	var params []interface{}
	slice1 := append(params, 2)
	// price := "product_name"
	// prdtID := "product_id"
	query := "DELETE FROM product WHERE product_id = ?"
	err = provideDB.Update(query, slice1)
	if err != nil {
		log.Println("unable to delete", err)
	}

	log.Println("updated")
}
