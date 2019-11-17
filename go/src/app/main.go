package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        a := c.DefaultQuery("go", "Guest")
        c.JSON(200, gin.H{
            "message": a,
        })
    })
    r.Run()
}
