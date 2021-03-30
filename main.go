package main

import (
	"context"
	"log"
	"os"

	"github.com/mmuoDev/commons/mongo"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson"
	//"encoding/json"
)

// "fmt"

// "context"
// "fmt"
// "log"
// "os"

type Partner struct {
	ID              *string `bson:"id,omitempty"`
	Name            *string `bson:"name,omitempty"`
	Address         *string `bson:"address,omitempty"`
	MultitenancyKey *string `bson:"multitenancykey,omitempty"`
}

func main() {
	//port := "9000"
	os.Setenv("MONGO_URL", "mongodb://localhost:27017")
	os.Setenv("MONGO_DB_NAME", "rrs")
	provide, err := mongo.NewConfigFromEnvVars().ToProvider(context.Background())
	if err != nil {
		log.Fatal("unable to connect mongo")
	}
	log.Println("connected to DB")
	col := mongo.NewCollection(provide, "mare_accounts")

	query := bson.M{
		"id":             "a139ec0f-1c80-414b-99b0-d121beb5b106",
		"owners.tenantId": "dcf",
	}

	update := bson.M{
		"owners.$.phoneNumber": "+2348034741353"}

	if err := col.UpdateWithFilterOptions(query, update); err != nil {
		log.Fatal(err)
	}

	log.Println("done!")
	// str := `{"id": "1234","name": "rocks", "address":"lagos", "multitenancykey":""}`
	// var res Partner
	// err1 := json.Unmarshal([]byte(str), &res)
	// if err1 != nil {
	// 	log.Panic(err1)
	// }
	// var partner Partner

	// mtk := ""
	// s := Partner{Name: "uche", ID : "1234", Address: "New York", MultitenancyKey: &mtk}
	//update := Partner{Name: "uche", ID : "1234", Address: "New York", MultitenancyKey: "089636"}

	// err2 := col.UpdateOneWithResult("1234", res, &partner)
	// if err2 != nil {
	// 	log.Println("error")
	// }
	// log.Println(partner.Name)

}
