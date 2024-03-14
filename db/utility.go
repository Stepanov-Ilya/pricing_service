package db

import (
	"database/sql"
	"fmt"
	"purple_hack_tree/structures"
	"regexp"
	"sort"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var STORAGE CurrentStorage

func GetPrice(request structures.Request) structures.Response {
	discountIds := GetSegmentsByUserID(request.UserId)
	sort.Slice(discountIds, func(i, j int) bool { return discountIds[i] > discountIds[j] })
	var response structures.Response
	if discountIds != nil {
		//TODO
		//for _, discountId := range discountIds {
		//	// Todo search in storage of discount
		//	// FindInDiscount(&response, discountId, request.LocationId, request.MicroCategoryId)
		//	// if response != nil {
		//	//	return response
		//	//}
		//
		//}
	}

	// Todo search in storage of baseline
	//FindInBaseline(&response, request.LocationId, request.MicroCategoryId)

	return response
}

func UpdateStorage() {

	ClearStorageSegments()
	wg.Add(1)
	wg.Add(1)
	go UpdateBaseline()
	go UpdateDiscounts()

	wg.Wait()
}

func UpdateBaseline() {

	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	STORAGE.Baseline++
	tableNew := "baseline_matrix_" + strconv.FormatUint(STORAGE.Baseline, 10)

	query := `
        CREATE TABLE IF NOT EXISTS ` + tableNew + ` (
            id INT AUTO_INCREMENT PRIMARY KEY,
            microcategory INT,
	      	location INT,
	     	price INT  
        )
    `
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO ` + tableNew + ` (id, microcategory, location, price) SELECT id, microcategory, location, price FROM process_baseline`)
	if err != nil {
		panic(err)
	}

	defer wg.Done()
}

func UpdateDiscounts() {
	segments, err := GetAllSegments()
	if err != nil {
		print(err)
	}
	STORAGE.Discounts = make(map[uint64]uint64)
	for _, segment := range segments {
		wg.Add(1)
		UpdateDiscountsMatrix(segment)
		STORAGE.Discounts[segment] = STORAGE.MaxDiscount
	}

	defer wg.Done()
}

func AddProcessBaseline(microcategory uint64, location uint64, price uint64) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "INSERT INTO process_baseline (microcategory, location, price) VALUES (?, ?, ?)"
	_, err = db.Exec(query, microcategory, location, price)
	if err != nil {
		panic(err)
	}
	db.Close()
}

func AddProcessDiscounts(segment uint64, microcategory uint64, location uint64, price uint64) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var segmentExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM discount_segments WHERE segment = ?)", segment).Scan(&segmentExists)
	if err != nil {
		panic(err)
	}

	if !segmentExists {
		_, err = db.Exec("INSERT INTO discount_segments (segment) VALUES (?)", segment)
		if err != nil {
			panic(err)
		}
	}

	tableName := "segment_" + strconv.FormatUint(segment, 10) // Prepend with a letter

	exists, err := tableExists(db, tableName)
	if err != nil {
		print(err)
	}

	if !exists {
		query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INT AUTO_INCREMENT PRIMARY KEY,
			microcategory INT,
			location INT,
			price INT
		)`, tableName)
		_, err = db.Exec(query)
		if err != nil {
			panic(err)
		}

	}

	query := fmt.Sprintf(`INSERT INTO %s (microcategory, location, price) VALUES (?, ?, ?)`, tableName)
	_, err = db.Exec(query, microcategory, location, price)
	if err != nil {
		panic(err)
	}
}

func tableExists(db *sql.DB, tableName string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)`
	var exists bool
	err := db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func GetCurrentStorage() {
	baseline, err := GetBaselineTables()
	if err != nil {
		print(err)
	}
	discount, err := GetDiscountTables()
	if err != nil {
		print(err)
	}

	STORAGE.MaxDiscount = uint64(discount[0])
	STORAGE.Baseline = uint64(baseline[0])
	STORAGE.Discounts, err = GetSegmentData()
	if err != nil {
		print(err)
	}
}

func GetBaselineTables() ([]int, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tables []int
	rows, err := db.Query("SHOW TABLES LIKE 'baseline_matrix_%'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tableNumber, err := extractBaselineNumber(tableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tableNumber)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(tables)))

	return tables, nil
}

func GetDiscountTables() ([]int, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tables []int
	rows, err := db.Query("SHOW TABLES LIKE 'discount_matrix_%'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tableNumber, err := extractDiscountNumber(tableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tableNumber)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(tables)))

	return tables, nil
}

func GetSegmentData() (map[uint64]uint64, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT segment, bd FROM discount_segments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	segmentData := make(map[uint64]uint64)

	for rows.Next() {
		var segmentID int
		var bd int
		if err := rows.Scan(&segmentID, &bd); err != nil {
			return nil, err
		}
		segmentData[uint64(segmentID)] = uint64(bd)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return segmentData, nil
}

func extractBaselineNumber(tableName string) (int, error) {
	re := regexp.MustCompile(`baseline_matrix_(\d+)`)
	matches := re.FindStringSubmatch(tableName)
	if len(matches) < 2 {
		return 0, fmt.Errorf("unable to extract : %s", tableName)
	}
	tableNumber, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	return tableNumber, nil
}

func extractDiscountNumber(tableName string) (int, error) {
	re := regexp.MustCompile(`discount_matrix_(\d+)`)
	matches := re.FindStringSubmatch(tableName)
	if len(matches) < 2 {
		return 0, fmt.Errorf("unable to extract : %s", tableName)
	}
	tableNumber, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	return tableNumber, nil
}

func GetAllSegments() ([]uint64, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var segments []uint64

	rows, err := db.Query("SELECT segment FROM discount_segments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var segmentID int
		if err := rows.Scan(&segmentID); err != nil {
			return nil, err
		}
		segments = append(segments, uint64(segmentID))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return segments, nil
}

func UpdateDiscountsMatrix(segment uint64) {
	STORAGE.MaxDiscount++
	db, err := sql.Open("mysql", connection)
	if err != nil {
		print(err)
	}
	defer db.Close()

	tableNew := "discount_matrix_" + strconv.FormatUint(STORAGE.MaxDiscount, 10)

	query := `
        CREATE TABLE IF NOT EXISTS ` + tableNew + ` (
            id INT AUTO_INCREMENT PRIMARY KEY,
            microcategory INT,
	      	location INT,
	     	price INT  
        )
    `
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`INSERT INTO ` + tableNew + ` (id, microcategory, location, price) SELECT id, microcategory, location, price FROM segment_` + strconv.FormatUint(segment, 10))
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO storage_segments (segment, bd) VALUES (?, ?)", segment, STORAGE.MaxDiscount)
	if err != nil {
		print(err)
	}

	defer wg.Done()
}

func GetArrayOfBaseline() ([][]int64, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var data [][]int64

	rows, err := db.Query("SELECT microcategory, location, price FROM baseline_matrix_" + strconv.FormatUint(STORAGE.Baseline, 10))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var microcategory, location, price int64
		if err := rows.Scan(microcategory, location, price); err != nil {
			return nil, err
		}
		data = append(data, []int64{microcategory, location, price})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func GetArrayOfDiscount(dp uint64) ([][]int64, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var data [][]int64

	rows, err := db.Query("SELECT microcategory, location, price FROM dicount_matrix_" + strconv.FormatUint(dp, 10))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var microcategory, location, price int64
		if err := rows.Scan(microcategory, location, price); err != nil {
			return nil, err
		}
		data = append(data, []int64{microcategory, location, price})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func SelectBaseline(microcategory int64, location int64) (int64, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var price int64

	query := "SELECT price FROM baseline_matrix_" + strconv.FormatUint(STORAGE.Baseline, 10) + " WHERE microcategory = ? AND location = ?"

	err = db.QueryRow(query, microcategory, location).Scan(&price)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func SelectDiscount(segment int64, microcategory int64, location int64) (int64, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var price int64

	query := "SELECT price FROM discount_matrix_" + strconv.FormatUint(STORAGE.Discounts[uint64(segment)], 10) + " WHERE microcategory = ? AND location = ?"

	err = db.QueryRow(query, microcategory, location).Scan(&price)
	if err != nil {
		return 0, err
	}

	return price, nil
}
