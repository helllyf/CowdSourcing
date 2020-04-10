package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/helllyf/CrowdSourcing/eth"
	"log"
)

func initEth() string{
	client, err := eth.Connect("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	b, err := client.GetBlockNumber(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.String())
	return b.String()
}

func main() {
	num := initEth()
	r := gin.Default()
	r.Static("/pages", "./pages")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": num,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}