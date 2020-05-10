package handler

import (
	"log"
	"net/http"
	"database/sql"
	"server/src/model"
	"github.com/gin-gonic/gin"
)

func ReDirectTo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		short := c.Param("short")
		_ = short
		// Fetch this URL from DB
		// Check if same short URL exists or not in DB
		rows, err := db.Query(`SELECT * FROM url WHERE shortURL="` + short + `"` + ` AND ip="` + c.ClientIP() + `"`)
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

		// Allow to redirect only if its present in DB
		if len(urls) > 0 {
			redirectURL := urls[0]

			// Increment count for that URL
			redirectURL.Count = redirectURL.Count + 1

			// Update in DB
			stmt, err := db.Prepare(`UPDATE url SET count=? WHERE shortURL=?`)
			if err != nil {
				log.Fatal(err)
				c.Status(http.StatusInternalServerError)
			} 
			_, err = stmt.Exec(redirectURL.Count, short)
			if err != nil {
				log.Fatal(err)
				c.Status(http.StatusInternalServerError)
			} 

			// Redirect to Long URL for this short URL
			c.Redirect(301, redirectURL.LongURL)

		} else {
			c.JSON(http.StatusNoContent, gin.H{
				"success": false,
				"message": "No such short URL in the DB.",
			})
		}
	}
}
