package pgin

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/go-sql-driver/mysql"
	"time"
)

var sdb *sql.DB

func MySql(){
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root@123",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "mall",
	}
	// Get a database handle.
	var err error
	sdb, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	
	username := "jack"
	password := "secret"
	createdAt := time.Now()

	result, err := sdb.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
    if err !=nil {
		fmt.Println("插入错误:",err)
	}
	id, serr := result.LastInsertId()
	if serr != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("New record ID is:", id)
	


}

func CreateTable(){
	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

// Executes the SQL query in our database. Check err to ensure there was no error.
 _, err9 := sdb.Exec(query)
 if err9 != nil {
	fmt.Println("Connected!0000")
  }
}

func InsertData(){

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	result, err := sdb.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
    if err !=nil {
		fmt.Println("插入错误:",err)
	}
	id, serr := result.LastInsertId()
	if serr != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("New record ID is:", id)
}