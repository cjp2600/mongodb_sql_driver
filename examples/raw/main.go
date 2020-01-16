package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "mongodb-sql-driver"
)

func init() {
	_ = godotenv.Load(".env.raw")
}

func main() {

	db, err := sql.Open("mongodb", os.Getenv("MONGO_URL"))
	if err != nil {
		fmt.Print(err)
		return
	}

	// prepare two return objects
	rows := make([]*sql.Rows, 2)
	for i := 0; i != 2; i++ {
		rows[i], err = db.Query("SELECT * FROM orders")
		if err != nil {
			fmt.Print(err)
			return
		}
	}

	// read results
	for _, r := range rows {
		for r.Next() {
			var f1, f2, f3 string
			err := r.Scan(&f1, &f2, &f3)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Printf("first_name=%s, last_name=%s, username=%s\n", f1, f2, f3)
		}
	}

}
