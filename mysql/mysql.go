package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

//Insert inserts data into table. Expects a table name and a map of columns and its respective values and maybe last inserted ID is returned
func (dbProvider DbProviderFunc) Insert(tableName string, values map[string]interface{}, lastInsertID bool) (int64, int64, error) {
	columns := make([]string, 0, len(values))
	rows := []interface{}{}
	for k, v := range values {
		columns = append(columns, k)
		rows = append(rows, v)
	}

	cols := strings.Join(columns, ",")
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
		return 0, 0, errors.Wrap(err, "Unable to prepare statement")
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, rows...)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Error inserting data")
	}
	aRows, err := res.RowsAffected()
	if err != nil {
		return 0, 0, errors.Wrap(err, "Unable to find rows affected")
	}
	if lastInsertID {
		id, err := res.LastInsertId()
		if err != nil {
			return aRows, 0, errors.Wrap(err, "Unable to fetch last inserted ID")
		}
		return aRows, id, nil
	}
	return aRows, 0, nil
}

//
func GeneratePlaceHolder(len int) string {
	pH := make([]string, 0, len)
	for i := 1; i <= len; i++ {
		pH = append(pH, "?")
	}
	pHs := strings.Join(pH, ",")
	return "(" + pHs + ")"
}

//InsertMulti inserts multi data into table
func (dbProvider DbProviderFunc) InsertMulti(tableName string, values map[string][]interface{}) (int64, error) {
	var genPHs string
	vLen := len(values)
	genPHs = GeneratePlaceHolder(vLen)
	columns := make([]string, 0, vLen)
	var params []string
	for k, _ := range values {
		columns = append(columns, k)
		params = append(params, genPHs)
	}
	log.Println(strings.Join(params, ","))
	return 0, nil
	//TODO: figure out how the func arguments should look like.
}

//Select returns a single row
func (dbProvider DbProviderFunc) Select(query string, params []interface{}, cols ...interface{}) (interface{}, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := dbProvider().PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to prepare statement")
	}
	res := stmt.QueryRow(params...)
	if err1 := res.Scan(cols...); err1 != nil {
		return nil, errors.Wrap(err1, "Unable to fetch rows")
	}
	return cols, nil
	//case
	// var params []interface{}
	// slice1 := append(params, 3)
	// price := "product_name"
	// prdtID := "product_id"
	// query := "SELECT product_price, product_name FROM product WHERE product_id = ?"
	// _, err = provideDB.Select(query, slice1, &price, &prdtID)
	// if err != nil {
	// 	log.Println("new error", err)
	// }
	
	// log.Println(prdtID)
}

//SelectMulti returns more than one row
func (dbProvider DbProviderFunc) SelectMulti(query string, params []interface{}) (*sql.Rows, error) {
	
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := dbProvider().PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to prepare statement")
	}
	rows, err1 := stmt.Query(params...)
	if err1 != nil {
		return nil, errors.Wrap(err1, "Unable to fetch rows")
	}
	
	return rows, nil

	// type Product struct {
	// 	productID int
	// 	productName string
	// }
	// for rows.Next() {
	// 	product := Product{}
	// 	err = rows.Scan(&product.productID, &product.productName)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println(product)
    // }
}

//Update updates a table
func (dbProvider DbProviderFunc) Update(query string, params []interface{}) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := dbProvider().PrepareContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "Unable to prepare statement")
	}
	_, err1 := stmt.Exec(params...)
	if err1 != nil {
		return errors.Wrap(err1, "Unable to update table")
	}
	return nil
}

func (dbProvider DbProviderFunc) Delete(query string, params []interface{}) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := dbProvider().PrepareContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "Unable to prepare statement")
	}
	_, err1 := stmt.Exec(params...)
	if err1 != nil {
		return errors.Wrap(err1, "Unable to delete table")
	}
	return nil
}



