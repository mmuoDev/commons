package mongo

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DbConfig vars
type DbConfig struct {
	DBURL  string
	DBName string
}

// ToProvider returns a mongo db provider from the config
func (c DbConfig) ToProvider(ctx context.Context) (DbProviderFunc, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI(c.DBURL))
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to connect to mongo using config=%+v", c)
	}
	return DbProvider(mClient, c.DBName), nil
}

// NewConfigFromEnvVars returns mongo configuration from environment variables
func NewConfigFromEnvVars() DbConfig {
	return DbConfig{
		DBURL:  os.Getenv("MONGO_URL"),
		DBName: os.Getenv("MONGO_DB_NAME"),
	}
}
