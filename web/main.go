package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artist", handlers.Artist)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	fmt.Println("Listening on http://localhost:8080/ ... ")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error : cant open port in http://localhost:8080/ ... ")
		return
	}
}
