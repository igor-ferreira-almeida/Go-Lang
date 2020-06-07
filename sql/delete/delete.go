package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("delete usuarios set nome = ? where id = ?")

	stmt.Exec("Órion", 1)
	stmt.Exec("Abel", 2)

}
