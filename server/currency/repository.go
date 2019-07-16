package currency

import (
	"fmt"
	"trainingApi/server/currency/models"

	"github.com/jinzhu/gorm"
)

// Repository with all methods to interact with db
// type Repository struct {
// 	db *gorm.DB
// }

// // NewRepository constructor
// func NewRepository(db *gorm.DB) *Repository {
// 	return &Repository{
// 		db,
// 	}
// }

// FindByID is a repository function to interact through gorm to get a currency
func FindByID(id string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency
	if err := db.Where("id = ?", id).First(&currency).Error; err != nil {
		return currency, err
	}

	return currency, nil
}

// FindBestBuy is a repository funtion to find best buy value currency
func FindBestBuy(name string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency
	if err := db.Where("name = ?", name).Order("buy desc").First(&currency).Error; err != nil {
		return currency, err
	}

	return currency, nil
}

// FindBestSell is a repository funtion to find best sell value currency
func FindBestSell(name string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency
	if err := db.Where("name = ?", name).Order("sell asc").First(&currency).Error; err != nil {
		return currency, err
	}

	return currency, nil
}

// FindAll is a repository function to find all currencies
func FindAll(db *gorm.DB) ([]models.Currency, error) {
	var currencies []models.Currency
	if err := db.Find(&currencies).Error; err != nil {
		return currencies, err
	}

	return currencies, nil
}

// Repository something
// type Repository interface {
// 	Find(out interface{}, where ...interface{}) *gorm.DB
// }

// type MockDb struct {
// 	Error error
// }

// func (m *MockDb) Find(out interface{}, where ...interface{}) *gorm.DB {
// 	out =
// 	return m
// }

// FindAllBestSellBuy is a repository function to find the best buy/sell of all currencies
func FindAllBestSellBuy(db *gorm.DB) ([]models.Currency, []models.Currency, error) {
	var currencies []models.Currency
	if err := db.Find(&currencies).Error; err != nil {
		return nil, nil, err
	}

	type ReturnValues struct {
		BestBuy  []models.Currency
		BestSell []models.Currency
	}

	// Get best buy values
	buyCurrencies := make(map[string]models.Currency)

	for _, currency := range currencies {
		// var found bool
		if _, ok := buyCurrencies[currency.Name]; ok == false {
			buyCurrencies[currency.Name] = currency
		}

		if buyCurrencies[currency.Name].Buy < currency.Buy {
			buyCurrencies[currency.Name] = currency
		}
	}

	// Get best sell values
	sellCurrencies := make(map[string]models.Currency)

	for _, currency := range currencies {
		// var found bool
		if _, ok := sellCurrencies[currency.Name]; ok == false {
			sellCurrencies[currency.Name] = currency
		}

		if sellCurrencies[currency.Name].Buy > currency.Buy {
			sellCurrencies[currency.Name] = currency
		}
	}

	finalCurrencies := &ReturnValues{
		BestBuy:  make([]models.Currency, 0, len(buyCurrencies)),
		BestSell: make([]models.Currency, 0, len(sellCurrencies)),
	}

	for _, currency := range buyCurrencies {
		finalCurrencies.BestBuy = append(finalCurrencies.BestBuy, currency)
	}

	for _, currency := range sellCurrencies {
		finalCurrencies.BestSell = append(finalCurrencies.BestSell, currency)
	}

	fmt.Println("buy", finalCurrencies.BestBuy)
	fmt.Println("sell", finalCurrencies.BestSell)

	return finalCurrencies.BestBuy, finalCurrencies.BestSell, nil
}
