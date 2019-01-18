package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/pdf", pdf)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func pdf(c *gin.Context) {
	bytes, err := getPdf()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/pdf", bytes)
}

func getPdf() ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8081", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
