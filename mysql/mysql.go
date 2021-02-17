package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" //
	"github.com/pkg/errors"
)

//DbProviderFunc provides a mysql database
type DbProviderFunc func() *sql.DB

//DbProvider is mysql db provider
func DbProvider(c *sql.DB) DbProviderFunc {
	return func() *sql.DB {
		return c
	}
}

//PingConnnect pings a mysql connection
func (dbProvider DbProviderFunc) PingConnnect() (bool, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err := dbProvider().PingContext(ctx)
	if err != nil {
		return false, errors.Wrapf(err, "Unable to ping the DB")
	}
	return true, nil
}

//Create creates a table
func (dbProvider DbProviderFunc) Create(query string) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := dbProvider().ExecContext(ctx, query)
	if err != nil {
		return errors.Wrapf(err, "Unable to create table")
	}
	return nil
}

//Insert inserts data into table
func (dbProvider DbProviderFunc) Insert(tableName string, values map[string]interface{}) (int64, error){
	columns := make([]string, 0, len(values))
	rows := make([]string, 0, len(values))
	for k, v := range values {
		columns = append(columns, k)
		//string
		log.Printf("%T", v)
		s, ok := v.(string)
		if ok {
			rows = append(rows, s)
		}
		n, ok := v.(int)
		t := strconv.Itoa(n)
		if 	ok {
			rows = append(rows, t)
		}
		f, ok := v.(float64)
		if ok {
			sf := fmt.Sprintf("%f", f)
			rows = append(rows, sf)
		}
		
	}

	cols := strings.Join(columns, ",")
	rowsV := strings.Join(rows, ",")
	//placeholders
	mLen := len(columns)
	pH := make([]string, 0, mLen)
	for i := 1; i <= mLen; i++ {
		pH = append(pH, "?")
	}
	pHs := strings.Join(pH, ",")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", tableName, cols, pHs)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    stmt, err := dbProvider().PrepareContext(ctx, query)
    if err != nil {
        return 0, errors.Wrap(err, "Unable to prepare statement")
	}
	defer stmt.Close()
	log.Println("rows", rowsV)
	res, err := stmt.ExecContext(ctx, rowsV)
	if err != nil {
		return 0, errors.Wrap(err, "Error inserting data")
	}
	//rows affected
	aRows, err := res.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "Unable to find rows affected")
	}
	return aRows, nil

}
