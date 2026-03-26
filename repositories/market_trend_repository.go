package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type marketTrendRepository struct {
	db *gorm.DB
}

func NewMarketTrendRepository(db *gorm.DB) MarketTrendRepository {
	return &marketTrendRepository{db}
}

func (r *marketTrendRepository) Create(trend *models.MarketTrend) error {
	return r.db.Create(trend).Error
}

func (r *marketTrendRepository) FindAll() ([]models.MarketTrend, error) {
	var trends []models.MarketTrend
	err := r.db.Find(&trends).Error
	return trends, err
}

func (r *marketTrendRepository) FindByID(id uint) (*models.MarketTrend, error) {
	var trend models.MarketTrend
	err := r.db.First(&trend, id).Error
	if err != nil {
		return nil, err
	}
	return &trend, nil
}

func (r *marketTrendRepository) Update(trend *models.MarketTrend) error {
	return r.db.Save(trend).Error
}

func (r *marketTrendRepository) Delete(id uint) error {
	return r.db.Delete(&models.MarketTrend{}, id).Error
}

func (r *marketTrendRepository) FindTopTen() ([]models.MarketTrend, error) {
	var trends []models.MarketTrend
	err := r.db.Order("demand_score desc").Limit(10).Find(&trends).Error
	return trends, err
}
func (r *marketTrendRepository) ReplaceAll(trends []models.MarketTrend) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.MarketTrend{}).Error; err != nil {
			return err
		}
		if len(trends) > 0 {
			if err := tx.Create(&trends).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
