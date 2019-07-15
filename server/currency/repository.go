package currency

import (
	"trainingApi/server/currency/models"

	"github.com/jinzhu/gorm"
)

// Repository with all methods to interact with db
type Repository struct {
	db *gorm.DB
}

// NewRepository constructor
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db,
	}
}

// FindCurrency is a repository function to interact through gorm to get a currency
func (r *Repository) FindCurrency(id string) (models.Currency, error) {
	var currency models.Currency
	if err := r.db.Where("id = ?", id).First(&currency).Error; err != nil {
		return currency, err
	}

	return currency, nil
}

// FindBestBuyCurrency is a repository funtion to find best buy value currency
func (r *Repository) FindBestBuyCurrency(name string) (models.Currency, error) {
	var currency models.Currency
	if err := r.db.Where("name = ?", name).Order("buy desc").First(&currency).Error; err != nil {
		return currency, err
	}

	return currency, nil
}

// FindBestSellCurrency is a repository funtion to find best sell value currency
func (r *Repository) FindBestSellCurrency(name string) (models.Currency, error) {
	var currency models.Currency
	if err := r.db.Where("name = ?", name).Order("sell asc").First(&currency).Error; err != nil {
		return currency, err
	}

	return currency, nil
}

// FindCurrencies is a repository function to find get all currencies
func (r *Repository) FindCurrencies() ([]models.Currency, error) {
	var currencies []models.Currency
	if err := r.db.Find(&currencies).Error; err != nil {
		return currencies, err
	}

	return currencies, nil
}
