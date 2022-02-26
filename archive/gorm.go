package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // runs required init() code
)

type User struct {
	gorm.Model // composition
	Name       string
	Email      string `gorm:"not null;unique_index"`
	Orders     []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func gorm_main() {
	fmt.Println("Connecting to DB...")
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := gorm.Open("postgres", psqInfo)
	must(err)
	fmt.Println("Connected")

	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&User{}, &Order{}) // only happens once
	user := getFirstRow(db)

	// createOrder(db, user, 10, "Description 1")
	// createOrder(db, user, 20, "Description 2")
	// createOrder(db, user, 30, "Description 3")

	db.Preload("Orders").First(&user) // DB join
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println("Email:", user.Email)
	fmt.Println("Number of orders:", len(user.Orders))
	fmt.Println("Orders:", user.Orders)
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	})
	if db.Error != nil {
		panic(db.Error)
	}
}

func insertToDb(db *gorm.DB) {
	name, email := getInfo()
	u := &User{
		Name:  name,
		Email: email,
	}
	if err := db.Create(u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)
}

func getFirstRow(db *gorm.DB) User {
	var u User
	id := 1
	db.First(&u, id) // adds a LIMIT 1 in the SQL query
	must(db.Error)

	fmt.Println(u)

	maxId := 3
	db.Where("id <= ?", maxId).First(&u)
	must(db.Error)

	fmt.Println(u)

	u.Email = "jon@calhoun.io"
	db.Where(u).First(&u)
	must(db.Error)

	fmt.Println(u)
	return u
}

func getManyRows(db *gorm.DB) {
	var users []User
	db.Find(&users)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println("Retrieved", len(users), "users.")
	fmt.Println(users)
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name:")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Enter email:")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}
