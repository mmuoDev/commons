package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql" //
)

//DbConfig vars
type DbConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBName     string
}

// NewConfigFromEnvVars returns mysql configuration from environment variables
func NewConfigFromEnvVars() DbConfig {
	return DbConfig{
		DBUsername: os.Getenv("MYSQL_USERNAME"),
		DBPassword: os.Getenv("MYSQL_PASSWORD"),
		DBHost:     os.Getenv("MYSQL_HOST"),
		DBName:     os.Getenv("MYSQL_DB_NAME"),
	}
}

//ConnectStr returns connection string
func (c DbConfig) ConnectStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.DBUsername, c.DBPassword, c.DBHost, c.DBName)
}

// ToProvider returns a mysql provider from the config
func (c DbConfig) ToProvider() (DbProviderFunc, error) {
	db, err := sql.Open("mysql", c.ConnectStr())
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to connect to mongo using config=%+v", c)
	}
	return DbProvider(db), nil
}
