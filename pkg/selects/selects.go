package selects

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

func GetAll(dbConn *sqlx.DB, tableName string, limit int) ([]map[string]interface{}, error) {

	sql := `SELECT * FROM ` + tableName + ` ORDER BY id`
	if limit > 0 {
		sql = sql + ` LIMIT ` + strconv.Itoa(limit)
	}

	fmt.Printf(sql + "\n")
	rows, err := dbConn.Query(sql)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return nil, err
	}

	count := len(cols)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePointers := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePointers[i] = &values[i]
		}
		err := rows.Scan(valuePointers...)
		if err != nil {
			fmt.Println("Failed to scan values", err)
			return nil, err
		}
		entry := make(map[string]interface{})
		for i, col := range cols {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	return tableData, nil
}

func GetById(dbConn *sqlx.DB, tableName string, entityId int) ([]map[string]interface{}, error) {

	sql := `SELECT * FROM ` + tableName + ` WHERE id=` + strconv.Itoa(entityId)

	fmt.Printf(sql + "\n")
	rows, err := dbConn.Query(sql)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return nil, err
	}

	count := len(cols)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePointers := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePointers[i] = &values[i]
		}
		err := rows.Scan(valuePointers...)
		if err != nil {
			fmt.Println("Failed to scan values", err)
			return nil, err
		}
		entry := make(map[string]interface{})
		for i, col := range cols {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	return tableData, nil
}
