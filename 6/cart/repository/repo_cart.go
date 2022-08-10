package repository

import (
	"context"
	"technical-test-telkom/cart"
	"technical-test-telkom/helper/logger"
	"technical-test-telkom/models"

	"gorm.io/gorm"
)

type cartRepository struct {
	Conn *gorm.DB
	log  logger.Logger
}



// NewSampleModuleRepository will create an object that represent the sample_module.Repository interface
func NewCartRepository(Conn *gorm.DB, log logger.Logger) cart.Repository {
	return &cartRepository{Conn, log}
}


func (c cartRepository) Find(ctx context.Context) ([]models.CartProduk, error) {
	var result []models.CartProduk
	err := c.Conn.WithContext(ctx).Find(&result).Error
	if err != nil{
		return nil, err
	}
	return result,nil
}

func (c cartRepository) GetByKodeProduk(ctx context.Context, kodeProduk string) (*models.CartProduk, error) {
	var result *models.CartProduk
	err := c.Conn.WithContext(ctx).First(&result,"kode_produk = ?",kodeProduk).Error
	if err != nil{
		return nil, err
	}
	return result,nil
}

func (c cartRepository) Delete(ctx context.Context, kodeProduk string) error {
	return c.Conn.WithContext(ctx).Where("kode_produk = ?",kodeProduk).Delete(&models.CartProduk{}).Error
}

func (c cartRepository) Store(ctx context.Context, insert models.CartProduk) error {
	return c.Conn.WithContext(ctx).Create(insert).Error
}

func (c cartRepository) Update(ctx context.Context, update models.CartProduk) error {
	return c.Conn.WithContext(ctx).Where("kode_produk = ?",update.KodeProduk).Updates(update).Error
}