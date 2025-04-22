package server

import (
	"fmt"
	"log"
	"os"
	"pt-xyz/configs/database"
	deliveryHttp "pt-xyz/internal/delivery/http"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/repository"
	"pt-xyz/internal/usecases"
	"pt-xyz/middlewares"

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
	
	adminRepository := repository.NewRepositoryAdmin(database.DB)
	consumerRepo := repository.NewRepositoryConsumer(database.DB)
	transactionRepo := repository.NewRepositoryTransaction()
	transactionProductRepo := repository.NewRepositoryTransactionProduct()
	loanLimitRepo := repository.NewRepositoryLoanLimit()
	loanInstallmentRepo := repository.NewRepositoryLoanInstallment()
	masterProductXyzRepo := repository.NewRepositoryMasterProductXYZ()

	adminService := usecases.NewServiceAdmin(adminRepository)
	consumerService := usecases.NewServiceConsumer(consumerRepo, adminRepository)
	transactionService := usecases.NewServiceTransaction(database.DB, transactionRepo, transactionProductRepo, loanLimitRepo, loanInstallmentRepo, masterProductXyzRepo)

	consumerHandler := deliveryHttp.NewHandlerConsumer(consumerService)
	transactionHandler := deliveryHttp.NewHandlerTransaction(transactionService)

	// Sensitive Data Exposure
	if os.Getenv("ADMIN_USERNAME") != "" && os.Getenv("ADMIN_PASSWORD") != "" && os.Getenv("ADMIN_FULLNAME") != "" {
        
        admin := entities.Admin{
            UserName: os.Getenv("ADMIN_USERNAME"),
            Password: os.Getenv("ADMIN_PASSWORD"),
            FullName: os.Getenv("ADMIN_FULLNAME"),
        }


        adminService.AddAdmin(&admin)
    } else {
        log.Println("Environment variables for ADMIN_USERNAME, ADMIN_PASSWORD, or ADMIN_FULLNAME are missing!")
    }

	

	router := gin.Default()
	router.Use(CORSMiddleware())

	v1 := router.Group("/v1")
	routes := v1.Group("") 
	
	routes.POST("/register", consumerHandler.Register)
	routes.POST("/login", consumerHandler.Login)
	
	// Broken Access control

	routes.POST("/transaction", middlewares.AuthMiddleware(), transactionHandler.CreateTransaction)
	


	routes.GET("/test-admin", middlewares.AuthMiddleware(), middlewares.AuthorizationMiddleware("admin"), func(c *gin.Context) {
		claims, _ := c.Get("claims")
		c.JSON(200, gin.H{
			"message": "Application Running",
			"data": claims,
		})
	})

	routes.GET("/test", middlewares.AuthMiddleware(), func(c *gin.Context) {
		claims, _ := c.Get("claims")
		c.JSON(200, gin.H{
			"message": "Application Running",
			"data": claims,
		})
	})

	return router.Run(":" + os.Getenv("PORT"))
}
