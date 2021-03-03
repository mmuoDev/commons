package main

import (

	// "fmt"

	"net/http"

	"github.com/mmuoDev/commons/httputils"
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

	var w http.ResponseWriter
	httputils.Get("https://jsonplaceholder.typicode.com/todos/1", w)
}
