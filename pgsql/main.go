package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// https://blog.csdn.net/u013210620/article/details/82702114
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "db_name"
)

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func query(db *sql.DB) {
	var id, name string

	rows, err := db.Query(" select * from hb where name=$1", 32)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id, name)
}

func main() {
	db := connectDB()
	query(db)

}
