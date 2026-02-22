package main
import ("fmt")

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default()
  fmt.Println("The value of router %v:",router)
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "Ping Pong",
    })
  })
  router.Run() // listens on 0.0.0.0:8080 by default
}