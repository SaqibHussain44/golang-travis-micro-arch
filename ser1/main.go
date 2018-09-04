package main

import (
	"log"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("init ser1")
}

func main() {
        http.HandleFunc("/A", Handler)
        log.Fatal(http.ListenAndServe(":8083", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    w.Write([]byte(`{"msg":"Welcome to service 1..!"}`))    
}

func GetTemp() int {
	return 100
}