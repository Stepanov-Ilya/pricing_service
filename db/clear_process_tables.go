package db

import (
	"database/sql"
	"fmt"
)

func ClearBaseline() {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM process_baseline")
	if err != nil {
		panic(err)
	}
}

func ClearDiscounts() {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT segment FROM discount_segments")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var segment uint64
		if err := rows.Scan(&segment); err != nil {
			panic(err)
		}

		tableName := fmt.Sprintf("segment_%d", segment)

		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
		if err != nil {
			panic(err)
		}
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		panic(err)
	}
}

func ClearDiscountsSegments() {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM discount_segments")
	if err != nil {
		print(err)
	}
}

func ClearStorageSegments() {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM storage_segments")
	if err != nil {
		print(err)
	}
}
