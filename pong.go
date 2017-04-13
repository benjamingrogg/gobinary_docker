package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
)

func main() {


	name, _ := os.Hostname()
	fmt.Println("Hostname : %s\n", name)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong", "hostname": name,
		})
	})
	r.Run(":8888") 
}
