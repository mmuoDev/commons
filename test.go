//port := "9000"
	//mongo
	os.Setenv("MONGO_URL", "mongodb://localhost:27017")
	os.Setenv("MONGO_DB_NAME", "esusu")
	provide, err := mongo.NewConfigFromEnvVars().ToProvider(context.Background())
	if err != nil {
		log.Fatal("unable to connect mongo")
	}
	log.Println("connected to DB")
	col := mongo.NewCollection(provide, "partners")
	//var partner Partner
	//update 
	update := Partner{Name: "jon-rose", Address:"colorado"}

	if err := col.Replace("1233456", &update); err != nil {
		log.Println("can't update", err)
	}
	log.Println(update)

	// query := "CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text, product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)"

	// err1 := provideDB.Create(query)
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }
	// log.Println("created!")
	os.Setenv("MYSQL_USERNAME", "root")
	os.Setenv("MYSQL_PASSWORD", "@Password12")
	os.Setenv("MYSQL_HOST", "127.0.0.1:3306")
	os.Setenv("MYSQL_DB_NAME", "pangaea")
	provideDB, err := mysql.NewConfigFromEnvVars().ToProvider()
	if err != nil {
		log.Fatal(err)
	}

	var params []interface{}
	//slice1 := append(params, 3)
	// price := "product_name"
	// prdtID := "product_id"
	query := "SELECT product_price, product_name FROM product"
	_, err = provideDB.SelectMulti(query, params)
	if err != nil {
		log.Println("new error", err)
	}

	log.Println("fetched")

	myMap := map[string]string{"name": "uche"}
	for key, value := range myMap {
		log.Println(key, value)
	}