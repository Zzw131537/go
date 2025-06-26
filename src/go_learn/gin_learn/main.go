package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Birthday string `json:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "123456")

		c.Next()

		latency := time.Since(t)
		log.Println(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}
func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/someJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	router.GET("/testing", startPage)

	router.Use(MyLogger())
	router.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println(example)
		c.JSON(200, gin.H{
			"example": example,
		})
	})
	router.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person

	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	c.String(http.StatusOK, "Success")
}
