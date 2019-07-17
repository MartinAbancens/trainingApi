package repository

import (
	model "trainingApi/server/models"

	"github.com/jinzhu/gorm"
)

// Repository is an interface for repositories
type Repository interface {
	GetByID(id string) (*model.Currency, error)
	GetByOrder(name string, order string) (*model.Currency, error)
	GetAll() (*[]model.Currency, error)
	GetBestBuySell() (*[]model.Currency, *[]model.Currency, error)
}

type repo struct {
	DB *gorm.DB
}

// CreateRepository is a function to initialize a repository
func CreateRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

// GetByID is a repository function to interact through gorm to get a currency
func (r *repo) GetByID(id string) (*model.Currency, error) {
	var currency model.Currency
	if err := r.DB.Where("id = ?", id).First(&currency).Error; err != nil {
		return &currency, err
	}

	return &currency, nil
}

// GetByOrder is a repository funtion to find best buy value currency
func (r *repo) GetByOrder(name string, order string) (*model.Currency, error) {
	var currency model.Currency
	if err := r.DB.Where("name = ?", name).Order(order).First(&currency).Error; err != nil {
		return &currency, err
	}

	return &currency, nil
}

// FindBestSell is a repository funtion to find best sell value currency
// func (r *repo) FindBestSell(name string) (model.Currency, error) {
// 	var currency model.Currency
// 	if err := db.Where("name = ?", name).Order("sell asc").First(&currency).Error; err != nil {
// 		return currency, err
// 	}

// 	return currency, nil
// }

// GetAll is a repository function to find all currencies
func (r *repo) GetAll() (*[]model.Currency, error) {
	var currencies []model.Currency
	if err := r.DB.Find(&currencies).Error; err != nil {
		return &currencies, err
	}

	return &currencies, nil
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

// GetBestBuySell is a repository function to find the best buy/sell of all currencies
func (r *repo) GetBestBuySell() (*[]model.Currency, *[]model.Currency, error) {
	var currencies []model.Currency
	if err := r.DB.Find(&currencies).Error; err != nil {
		return nil, nil, err
	}

	// Get best buy values
	buyCurrencies := make(map[string]model.Currency)

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
	sellCurrencies := make(map[string]model.Currency)

	for _, currency := range currencies {
		// var found bool
		if _, ok := sellCurrencies[currency.Name]; ok == false {
			sellCurrencies[currency.Name] = currency
		}

		if sellCurrencies[currency.Name].Buy > currency.Buy {
			sellCurrencies[currency.Name] = currency
		}
	}

	BestBuy := make([]model.Currency, 0, len(buyCurrencies))

	for _, currency := range buyCurrencies {
		BestBuy = append(BestBuy, currency)
	}

	BestSell := make([]model.Currency, 0, len(sellCurrencies))
	for _, currency := range sellCurrencies {
		BestSell = append(BestSell, currency)
	}

	return &BestBuy, &BestSell, nil
}
