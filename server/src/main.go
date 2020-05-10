package main

import (
	"fmt"
	"log"
	"os"
	"server/src/handler"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
		fmt.Println("Used CORS middleware!")
        c.Next()
    }
}

func main() {
	// Getting password from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	password := os.Getenv("PASSWORD")
	// Open port on AWS 
	db, err := sql.Open("mysql", "username:" + password + "password@tcp(54.162.135.16:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Testing the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Creating the DB
	// _, err = db.Exec("CREATE DATABASE URLShortener")
	// if err != nil {
	// 	log.Fatal(err)
	// } 
	
	// Choosing the DB
	_, err = db.Exec("USE URLShortener")
	if err != nil {
		log.Fatal(err)
	} 
	
	// Creating the table
	// stmt, err := db.Prepare("CREATE TABLE URL (short VARCHAR(20), long VARCHAR(100), count int, ip VARCHAR(20));")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = stmt.Exec()
	// if err != nil {
	// 	log.Fatal(err)
	// } 
	
	r := gin.Default()
	// Adding CORS middleware
	r.Use(CORSMiddleware())

	r.GET("/url", handler.URLGet(db))
	r.POST("/url", handler.URLPost(db))
	r.GET("/url/:short", handler.ReDirectTo(db))

	r.Run()
}
