package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/logins", Login)
	http.HandleFunc("/Home", Home)
	http.HandleFunc("/Refresh", Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
