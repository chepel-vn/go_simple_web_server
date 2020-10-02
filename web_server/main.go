package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Employee struct {
	Name   string `json:"empname"`
	Number int    `json:"empid"`
}

func handler(w http.ResponseWriter, req *http.Request) {
	emp := &Employee{Name: "Иванов Иван Иванович", Number: 1}
	js, err := json.Marshal(emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe("localhost:80", nil))
}
