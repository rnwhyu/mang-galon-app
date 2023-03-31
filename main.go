package main

import (
	"database/sql"
	"fmt"
	"galon-app/database"
	"galon-app/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	//config env
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file env.")
	} else {
		fmt.Println("success read file env.")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	// os.Getenv("PGHOST"),
	// 	os.Getenv("PGPORT"),
	// 	os.Getenv("PGUSER"),
	// 	os.Getenv("PGPASSWORD"),
	// 	os.Getenv("PGDATABASE"))
	fmt.Println(psqlInfo)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to DB")
	}
	database.DbMigrate(DB)
	defer DB.Close()

	r := gin.Default()
	router.Setup(r)
}
