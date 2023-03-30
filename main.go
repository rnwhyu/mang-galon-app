package main

import (
	"database/sql"
	"fmt"
	"galon-app/controller"
	"galon-app/controller/galon"
	"galon-app/database"
	"galon-app/middlewares"
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
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	protected := r.Group("/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controller.CurrentUser)
	protected.POST("/user/galon", galon.Add)
	protected.PUT("/user/galon/:id", galon.Update)
	protected.DELETE("/user/galon/:id", galon.Delete)
	protected.GET("/user/galon", galon.GetAll)

	r.Run(":" + os.Getenv("PORT"))
}
