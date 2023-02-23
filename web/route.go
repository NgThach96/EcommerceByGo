package main 

import (
    "github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello World",
    })
}
