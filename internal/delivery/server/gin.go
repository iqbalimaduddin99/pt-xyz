package server

import (
	"fmt"
	"log"
	"os"
	"pt-xyz/configs/database"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/repository"
	"pt-xyz/internal/usecases"

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
	
	if os.Getenv("ADMIN_USERNAME") != "" && os.Getenv("ADMIN_PASSWORD") != "" && os.Getenv("ADMIN_FULLNAME") != "" {
        
        admin := entities.Admin{
            UserName: os.Getenv("ADMIN_USERNAME"),
            Password: os.Getenv("ADMIN_PASSWORD"),
            FullName: os.Getenv("ADMIN_FULLNAME"),
        }

        adminRepository := repository.NewRepositoryAdmin(database.DB)
        adminService := usecases.NewServiceAdmin(adminRepository)

        adminService.AddAdmin(&admin)
    } else {
        log.Println("Environment variables for ADMIN_USERNAME, ADMIN_PASSWORD, or ADMIN_FULLNAME are missing!")
    }

	

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
