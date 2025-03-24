package server

import (
	"fmt"
	"os"
	"pt-xyz/configs/database"

	"github.com/gin-gonic/gin"
)


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func Run() error {
	err := database.Connect()
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer database.DB.Close()

	router := gin.Default()
	router.Use(CORSMiddleware())

	v1 := router.Group("/v1")
	 routes := v1.Group("") 

	 routes.GET("/test", func(c *gin.Context) {
		 c.JSON(200, gin.H{
			 "message": "Application Running",
		 })
	 })

	return router.Run(":" + os.Getenv("PORT"))
}
