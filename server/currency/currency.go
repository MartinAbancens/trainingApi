package currency

import (
	"trainingApi/server/currency/models"

	"github.com/jinzhu/gorm"
)

// FindCurrency is a repository function to interact through gorm to get a currency
func FindCurrency(id string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency
	if err := db.Where("id = ?", id).First(&currency).Error; err != nil {
		return currency, err
	}

	return currency, nil
}

// FindBestBuyCurrency is a repository funtion to find best buy value currency
func FindBestBuyCurrency(name string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency
	if err := db.Where("name = ?", name).Order("buy desc").First(&currency).Error; err != nil {
		return currency, err
	}

	return currency, nil
}

// FindBestSellCurrency is a repository funtion to find best sell value currency
func FindBestSellCurrency(name string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency
	if err := db.Where("name = ?", name).Order("sell asc").First(&currency).Error; err != nil {
		return currency, err
	}

	return currency, nil
}

// FindCurrencies is a repository function to find get all currencies
func FindCurrencies(db *gorm.DB) ([]models.Currency, error) {
	var currencies []models.Currency
	if err := db.Find(&currencies).Error; err != nil {
		return currencies, err
	}

	return currencies, nil
}
