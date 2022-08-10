package cart

import (
	"context"
	"technical-test-telkom/models"
)

type Repository interface {
	Find(ctx context.Context) ([]models.CartProduk, error)
	GetByKodeProduk(ctx context.Context,kodeProduk string)(*models.CartProduk,error)
	Delete(ctx context.Context,kodeProduk string)error
	Store(ctx context.Context,insert models.CartProduk)error
	Update(ctx context.Context,update models.CartProduk)error
}
