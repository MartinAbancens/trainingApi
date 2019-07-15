package handler

import (

	// Models
	//"trainingApi/server/models"

	// Gin and Gorm
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// Repositories
	r "trainingApi/server/currency"
)

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
// 		c.JSON(200, currency)
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
// 			c.JSON(200, gin.H{"message": msg})
// 		}
// 	}
// }

// // CreateCurrency handles POST request to create new Currency
// func CreateCurrency(db *gorm.DB) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		var currency models.Currency
// 		c.BindJSON(&currency)
// 		db.Create(&currency)
// 		c.JSON(200, currency)
// 	}
// }

// GetCurrencyByID handles GET one Currency by ID
func GetCurrencyByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		currency, err := r.FindCurrency(c.Params.ByName("id"), db)
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, currency)
		}
	}
}

// GetBestBuyValue handles GET one currency by name (best buy value from bank)
func GetBestBuyValue(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		currency, err := r.FindBestBuyCurrency(c.Params.ByName("id"), db)
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, currency)
		}
	}
}

// GetBestSellValue handles GET one currency by name (best sell value from bank)
func GetBestSellValue(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		currency, err := r.FindBestSellCurrency(c.Params.ByName("id"), db)
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(100, currency)
		}
	}
}

// GetAllCurrencies handle GET all currencies from the db
func GetAllCurrencies(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		currencies, err := r.FindCurrencies(db)
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(100, currencies)
		}
	}
}
