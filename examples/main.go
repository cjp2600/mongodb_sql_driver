package main

import (
	"database/sql"
)

func main() {

	db, err := sql.Open("mongodb", "mongoDSN=mongodb+srv://admin:P95tXKQPogntFaaY@pixel-stage-ameug.mongodb.net/test?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}

}
