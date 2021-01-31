package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	basePath, _ = os.Getwd()
	dataPath = fmt.Sprintf("%s/%s", basePath, "Soal_5/server.log")
)

type Body struct {
	Counter int `json:"counter"`
}

func main() {
	r := gin.Default()

	r.Use(Logger())

	r.POST("/", handlerPost)

	log.Fatal(r.Run())
}

func handlerPost(c *gin.Context) {
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		createFile(c)
		t := time.Now()
		ipAddr := c.ClientIP()
		var counter = new(Body)
		c.ShouldBind(counter)
		headerData := c.Request.Header.Get("X-RANDOM")
		data := fmt.Sprintf("[%s] Success: POST http://%s {\"Counter\": \"%d\", \"X-RANDOM\": \"%s\"}", t.Format(time.RFC3339), ipAddr, counter.Counter, headerData)
		WriteFile(data, c)
	}
}

func createFile(c *gin.Context) {
	_, err := os.Stat(dataPath)
	if os.IsNotExist(err) {
		file, err := os.Create(dataPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		defer file.Close()
	}
}

func WriteFile(data string, c *gin.Context) {
	file, err := os.OpenFile(dataPath, os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	defer file.Close()

	_, err = file.WriteString(data+"\n")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	err = file.Sync()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}