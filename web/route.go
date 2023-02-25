package main 

import (
    "github.com/gin-gonic/gin"
)

func route() {
	router := gin.Default()

    router.GET("/", helloWorld)

    router.Run(":3000")
}

func helloWorld(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello World",
    })
}
