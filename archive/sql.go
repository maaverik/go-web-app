package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // we need some code that its init() function runs
)

func test(db *sql.DB) {
	// test connection
	must(db.Ping())
	fmt.Println("Successfully connected")
}

func insert(db *sql.DB) {
	var id int                                                                       // to store row id
	row := db.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", // we use placeholders to prevent SQL injection
		"Jose", "jose@email.com")
	err := row.Scan(&id) // using Scan to get row id since db.Exec() doesn't return it for postgres
	must(err)

	// insert data
	for i := 1; i < 6; i++ {
		// Create some fake data
		userId := 1
		if i > 3 {
			userId = 2
		}
		amount := 1000 * i
		description := fmt.Sprintf("Item No.%d", i)
		err = db.QueryRow(`
				INSERT INTO orders (user_id, amount, description)
				VALUES ($1, $2, $3)
				RETURNING id`,
			userId, amount, description,
		).Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("Created an order with the ID:", id)
	}
}

func queryOne(db *sql.DB) {
	// query a row
	var id int // to store row id
	var name, email string
	row := db.QueryRow(`
			SELECT id, name, email
			FROM users
			WHERE id >= $1`, 1)
	err := row.Scan(&id, &name, &email)
	must(err)
	fmt.Println("ID:", id, "Name:", name, "Email:", email)
}

func queryMany(db *sql.DB) {
	// query many rows
	var id int // to store row id
	var name, email string
	rows, err := db.Query(`
		SELECT id, name, email
		FROM users
		WHERE ID >= $1`,
		1)
	must(err)
	for rows.Next() { // rows initially points to location before first row; Next() returns false when done
		// rows.Next() works like python generator, there's no way to go back to a processed row
		rows.Scan(&id, &name, &email)
		fmt.Println("ID:", id, "Name:", name, "Email:", email)
	}
}

func sql_main() {
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqInfo)
	must(err)

	// func to call

	db.Close()
}
