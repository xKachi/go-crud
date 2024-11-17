package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB


func InitDB() {
	godotenv.Load()
	password := os.Getenv("DB_PASSWORD")

	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("root:%s@tcp(localhost:3306)/%s", password, "go_crud"))
	fmt.Println(DB)

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	UsersTableSchema := `
	CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(50) NOT NULL,
  		email VARCHAR(100) NOT NULL UNIQUE,
	)
	`

	_, err := DB.Exec(UsersTableSchema)

	if err != nil {
		panic("Could not create users table.")
	}

}
