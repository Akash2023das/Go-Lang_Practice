package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostPage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	address := c.Query("address")

	c.JSON(200, gin.H{
		"Name is = ":    name,
		"Age is = ":     age,
		"Address is = ": address,
	})

}

func main() {
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostPage)
	r.GET("/query", QueryString)

	r.Run()
}
