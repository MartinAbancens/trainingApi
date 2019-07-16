package handler

import (

	// Models
	//"trainingApi/server/models"

	// Gin and Gorm
	"fmt"
	"net/http"
	model "trainingApi/server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// Repositories
	repo "trainingApi/server/repository"
)

// var (
// 	findCurrency       = repo.GetByID
// 	findBestCurrency   = repo.GetByOrder
// 	findAllCurrencies  = repo.GetAll
// 	findAllBestSellBuy = repo.GetBestBuySell
// )

// // UpdateCurrency handles PUT to update a Currency
// func UpdateCurrency(db *gorm.DB) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		var currency models.Currency
// 		id := c.Params.ByName("id")
// 		if err := db.Where("id = ?", id).First(&currency).Error; err != nil {
// 			c.AbortWithStatus(404)
// 			fmt.Println(err)
// 		}
// 		c.BindJSON(&currency)
// 		db.Save(&currency)
// 		c.JSON(http.StatusOK, currency)
// 	}
// }

// // DeleteCurrency handles DELETE request
// func DeleteCurrency(db *gorm.DB) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		var currency models.Currency
// 		id := c.Params.ByName("id")
// 		fmt.Printf("The id requested is %s \n", id)
// 		if err := db.Where("id = ?", id).Delete(&currency).Error; err != nil {
// 			c.AbortWithStatus(400)
// 			fmt.Println(err)
// 		} else {
// 			msg := fmt.Sprintf("blog post %s has been deleted", id)
// 			c.JSON(http.StatusOK, gin.H{"message": msg})
// 		}
// 	}
// }

// // CreateCurrency handles POST request to create new Currency
// func CreateCurrency(db *gorm.DB) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		var currency models.Currency
// 		c.BindJSON(&currency)
// 		db.Create(&currency)
// 		c.JSON(http.StatusOK, currency)
// 	}
// }

// GetCurrencyByID handles GET one Currency by ID
func GetCurrencyByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		r := repo.CreateRepository(db)
		result, err := r.GetByID(c.Params.ByName("id"))
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// GetBestBuyValue handles GET one currency by name (best buy value from bank)
func GetBestBuyValue(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		r := repo.CreateRepository(db)
		result, err := r.GetByOrder(c.Params.ByName("id"), "buy desc")
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// GetBestSellValue handles GET one currency by name (best sell value from bank)
func GetBestSellValue(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		r := repo.CreateRepository(db)
		result, err := r.GetByOrder(c.Params.ByName("id"), "sell asc")
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// GetAllCurrencies handle GET all currencies from the db
func GetAllCurrencies(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		r := repo.CreateRepository(db)
		results, err := r.GetAll()
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(http.StatusOK, results)
		}
	}
}

// GetAllCurrenciesBestBuySell handle GET all the best prices of all currencies from the db
func GetAllCurrenciesBestBuySell(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		r := repo.CreateRepository(db)
		bestBuy, bestSell, err := r.GetBestBuySell()
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			buyAndSell := make(map[string]*[]model.Currency)

			buyAndSell["buy"] = bestBuy
			buyAndSell["sell"] = bestSell

			fmt.Println("final:", buyAndSell)
			c.JSON(http.StatusOK, buyAndSell)
		}
	}
}