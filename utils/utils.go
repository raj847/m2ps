package utils

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func ToString(i interface{}) string {
	log, _ := json.Marshal(i)
	logString := string(log)

	return logString
}

func IsConnected() bool {

	_, err := http.Get("https://www.google.com")
	if err != nil {
		return false
	}
	//_, err := http.Get("http://clients3.google.com/generate_204")
	//if err != nil {
	//	return false
	//}
	return true
}

func DBTransaction(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Rollback Panic
		} else if err != nil {
			tx.Rollback() // err is not nill
		} else {
			err = tx.Commit() // err is nil
		}
	}()
	err = txFunc(tx)
	return err
}
