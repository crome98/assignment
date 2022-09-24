package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InsertDevice(devices map[string]interface{}) bool {
	var stmt = "INSERT INTO device (id, type, latitude, longitude, status, timezone) VALUES (?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE (type, latitude, longitude, status, timezone) VALUES (?, ?, ?, ?, ?)"
	db, err := sql.Open("mysql", "awair:awair@/awair")
	if err != nil {
		fmt.Println("Cannot connect to DB")
		return false
	}
	for device := range devices {
		db.Exec(stmt)
	}
	return false
}
