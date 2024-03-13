package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	connection = "root:A12345678a@tcp(127.0.0.1:3306)/purple_hack"
)

func CreateProcessBd() {

	// Открытие соединения с базой данных PostgreSQL
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	  CREATE TABLE IF NOT EXISTS process_baseline (
	      id INT AUTO_INCREMENT PRIMARY KEY,
	      microcategory INT,
	      location INT,
	      price INT     
	  )
	`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
	  CREATE TABLE IF NOT EXISTS baseline_matrix_1 (
	      id INT AUTO_INCREMENT PRIMARY KEY,
	      microcategory INT,
	      location INT,
	      price INT     
	  )
	`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
	  CREATE TABLE IF NOT EXISTS discount_matrix_1 (
	      id INT AUTO_INCREMENT PRIMARY KEY,
	      microcategory INT,
	      location INT,
	      price INT     
	  )
	`)
	if err != nil {
		panic(err)
	}
	//_, err = db.Exec(`
	//  CREATE TABLE IF NOT EXISTS process_discounts (
	//      id INT AUTO_INCREMENT PRIMARY KEY,
	//      segment INT,
	//      microcategory INT,
	//      location INT,
	//      price INT
	//  )
	//`)
	//if err != nil {
	//	panic(err)
	//}

	_, err = db.Exec(`
	 CREATE TABLE IF NOT EXISTS discount_segments (
	     id INT AUTO_INCREMENT PRIMARY KEY,
	     segment INT
	 )
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS storage_segments (
		id INT AUTO_INCREMENT PRIMARY KEY,
	    segment INT,
		bd INT
	 )
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database tables created successfully")
}
