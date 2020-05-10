package handler

import (
	"log"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"server/src/model"
)

func URLPost(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.URL{}
		err := c.Bind(&requestBody)
		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusInternalServerError)
		} 
		temp := strings.Split(requestBody.ShortURL, "/")
		temp2 := temp[len(temp) - 1]
		// Check if same short URL exists or not in DB
		rows, err := db.Query(`SELECT shortURL FROM url WHERE shortURL="` + temp2 + `"`)
		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusInternalServerError)
		} 
		count := 0
		for rows.Next() {
			count++
		}

		// Not a unique short URL
		if count > 0 {
			c.JSON(http.StatusNoContent, gin.H{
				"success": false,
				"message": "Not a unique short URL.",
			})
		} else {
			// If not, then add to DB
			stmt, err := db.Prepare(`
				INSERT INTO url (shortURL, longURL, count, ip)
				VALUES(?, ?, ?, ?)
			`)
			if err != nil {
				log.Fatal(err)
				c.Status(http.StatusInternalServerError)
			}

			_, err = stmt.Exec(temp2, requestBody.LongURL, 0, c.ClientIP())
			if err != nil {
				log.Fatal(err)
				c.Status(http.StatusInternalServerError)
			}

			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})	
		}
	}
}
