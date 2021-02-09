package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

//DbProviderFunc provides a mongo database
type DbProviderFunc func() *mongo.Database

//DbProvider is mongo db provider
func DbProvider(c *mongo.Client, dbName string) DbProviderFunc {
	return func() *mongo.Database {
		return c.Database(dbName)
	}
}

//Collection is a representation of a mongo collection
type Collection struct {
	col *mongo.Collection
}

//NewCollection creates a new collection
func NewCollection(provideDB DbProviderFunc, colName string) Collection {
	col := provideDB().Collection(colName)
	return Collection{col: col}
}
