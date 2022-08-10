package usecase

import (
	"context"
	"gorm.io/gorm"
	"technical-test-telkom/cart"
	"technical-test-telkom/helper/logger"
	"technical-test-telkom/models"
	"time"
)

type cartUsecase struct {
	cartRepo       cart.Repository
	contextTimeout time.Duration
	log            logger.Logger
}

// NewSampleModuleUsecase will create new an SampleModuleUsecase object representation of sample_module.Usecase interface
func NewCartUsecase(cartRepo cart.Repository, log logger.Logger, timeout time.Duration) cart.UseCase {
	return &cartUsecase{
		cartRepo:       cartRepo,
		contextTimeout: timeout,
		log:            log,
	}
}

func (c cartUsecase) Find(ctx context.Context) ([]models.CartProduk, error) {
	return c.cartRepo.Find(ctx)
}

func (c cartUsecase) Delete(ctx context.Context, kodeProduk string) error {
	return c.cartRepo.Delete(ctx,kodeProduk)
}

func (c cartUsecase) Store(ctx context.Context, insert models.CartProduk) error {
	chart ,err := c.cartRepo.GetByKodeProduk(ctx,insert.KodeProduk)
	if err != nil && err != gorm.ErrRecordNotFound{
		return err
	}
	if chart != nil{
		chart.Kuantitas = chart.Kuantitas + insert.Kuantitas
		return c.cartRepo.Update(ctx,*chart)
	}
	return c.cartRepo.Store(ctx,insert)
}