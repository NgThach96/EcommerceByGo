package main 

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", helloWorld)

    router.Run(":3000")
}
