package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// main function
func main() {
	//connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create table if not exists
	_, err = db.Exec("CREATE TABLE if NOT EXISTS users (id SERIAL PRIMARY KEY, NAME TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// create router
	router := mux.NewRouter()
	router.HandleFunc("/api/go/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/api/go/users", createUser(db)).Methods("POST")
	router.HandleFunc("/api/go/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/api/go/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/api/go/users/{id}", deleteUser(db)).Methods("DELETE")

	// wrap the router with CORS and JSON content type middlewares
	enhancedRouter := enableCORS(jsonContentTypeMiddleware(router))

	// start server
	log.Fatal(http.ListenAndServe(":8000", enhancedRouter))
}
