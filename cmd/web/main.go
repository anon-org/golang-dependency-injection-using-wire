package main

import (
	"database/sql"
	"example/user"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userHandler := user.Wire(db)
	http.Handle("/user", userHandler.FetchByUsername())
	http.ListenAndServe(":8000", nil)
}
