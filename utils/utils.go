package utils

import (
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
