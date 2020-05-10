package handler

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"server/src/model"
)

func URLGet(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		clientIP := c.ClientIP()
		rows, err := db.Query(`SELECT * FROM url WHERE ip="` + clientIP + `"`)
		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusInternalServerError)
		} 

		var urls []model.URL = []model.URL{}
		var url model.URL
		for rows.Next() {
			err = rows.Scan(&url.ShortURL, &url.LongURL, &url.Count, &url.IP)
			if err != nil {
				log.Fatal(err)
				c.Status(http.StatusInternalServerError)
			} 
			urls = append(urls, url)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": urls,
		})
	}
}
