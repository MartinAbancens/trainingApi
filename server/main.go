package main

import (
	"log"
	"net/http"

	// DB packages - external
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// Gin packages
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	// DB - internal

	h "trainingApi/server/handlers"
	// "tonic/server/utils/db/migrate"
)

var (
	currencyHandler h.Handler
)

var db *gorm.DB
var err error

func main() {

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// dbConnStr := os.Getenv("MYSQLSTR")

	// Set up DB connection
	db, err = gorm.Open("mysql", "root:heloraktia@/training?charset=utf8&parseTime=True&loc=Local")
	// db, err = gorm.Open("mysql", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// db.AutoMigrate(&models.Currency{})
	// migrate.Start(db)

	c := currencyHandler.InitializeHandler(db)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../client", true)))

	// API Route Groups
	api := router.Group("/api/")
	{
		// Currency Subgroup
		currency := api.Group("/currency/")
		{
			currency.GET("/", c.GetAllValuesOrdered())

			currency.GET("/:id", c.GetValueByID())

			currency.GET("/:id/buy", c.GetValueByOrderDesc())

			currency.GET("/:id/sell", c.GetValueByOrderAsc())

			// currency.PUT("/:id", c.UpdateCurrency())

			// currency.POST("/", c.CreateCurrency())

			// currency.DELETE("/:id", c.DeleteCurrency())
		}
		// Begin api base routes
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "base route reached",
			})
		})
		// more api routes go here
		// example:
		// api.GET("/specific/path/to/route/:id", someHandler)
	}
	// Start and run the server
	router.Run(":3000")
}

// // A standard handler function
// func someHandler(c *gin.Context) {
// 	c.Header("Content-Type", "application/json")
// 	c.JSON(http.StatusOK, JSON_content)
// }
