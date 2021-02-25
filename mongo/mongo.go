package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

//Insert adds a new document to the DB
func (c *Collection) Insert(doc interface{}) (*mongo.InsertOneResult, error) {
	return c.col.InsertOne(context.Background(), doc)
}

//InsertMulti adds multiple documents
func (c *Collection) InsertMulti(doc []interface{}) (*mongo.InsertManyResult, error) {
	return c.col.InsertMany(context.Background(), doc)
}

//FindByID finds a document by ID (primary key)
func (c *Collection) FindByID(ID string, r interface{}) error {
	filter := bson.D{{"id", ID}}
	err := c.col.FindOne(context.Background(), filter).Decode(r)
	if err != nil {
		return errors.Wrap(err, "Unable to find entity")
	}
	return err
}

//FindOne returns a document
func (c *Collection) FindOne(filter interface{}, r interface{}) error {
	err := c.col.FindOne(context.Background(), filter).Decode(r)
	if err != nil {
		return errors.Wrap(err, "unable to find entity")
	}
	return err
}

//FindMulti returns all document based on criteria
func (c *Collection) FindMulti(filter interface{}, onEach func(c *mongo.Cursor) error) error {
	findOptions := options.Find()
	ctx := context.Background()
	var cur *mongo.Cursor
	var err error

	if findOptions != nil {
		cur, err = c.col.Find(ctx, filter, findOptions)
	} else {
		cur, err = c.col.Find(ctx, filter)
	}

	if err != nil {
		return errors.Wrap(err, "unable to find documents")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		err := onEach(cur)
		if err != nil {
			return errors.Wrap(err, "unable to find documents")
		}
	}
	if err := cur.Err(); err != nil {
		return errors.Wrap(err, "unable to find documents")
	}
	return nil
	//sample
	// var pp []Partner
	// onEach := func(cur *mongo.Cursor) error {
	// 	p := Partner{}
	// 	if err := cur.Decode(&p); err != nil {
	// 		return err
	// 	}
	// 	pp = append(pp, p)
	// 	return nil
	// }

	// if err := col.FindMulti(filter, onEach); err != nil {
	// 	return error
	// }
}

//Replace replaces one document with another
func (c *Collection) Replace(ID string, replacement interface{}) error {
	filter := bson.D{{"id", string(ID)}}
	_, err := c.col.ReplaceOne(context.Background(), filter, replacement)
	return err
}

//Update updates an existing document in the database
func (c *Collection) Update(ID string, changes interface{}) error {
	update := bson.D{
		{"$set", changes},
	}
	filter := bson.D{{"id", string(ID)}}
	_, err := c.col.UpdateOne(context.Background(), filter, update)
	return err
}

//CountDocuments returns a count of all documents
func (c *Collection) CountDocuments(filter interface{}) (int64, error) {
	count, err := c.col.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return count, err
}

//UpdateMany updated one document
func (c *Collection) UpdateMany(filter interface{}, replacement interface{}) error {
	//filter - bson.D{{"id", "12345"}}
	//bson.D{{"$set", bson.D{{"author", "Nic Raboy"}}},}
	_, err := c.col.UpdateMany(context.Background(), filter, replacement)
	if err != nil {
		return errors.Wrap(err, "Unable to update")
	}
	return nil
}

//DeleteMany deletes document(s)
func (c *Collection) DeleteMany(filter interface{}) error {
	_, err := c.col.DeleteMany(context.Background(), filter)
	if err != nil {
		return errors.Wrap(err, "Unable to delete")
	}
	return nil
}
