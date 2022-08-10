package cart

import (
	"context"
	"technical-test-telkom/models"
)

type UseCase interface {
	Find(ctx context.Context) ([]models.CartProduk, error)
	Delete(ctx context.Context,kodeProduk string)error
	Store(ctx context.Context,insert models.CartProduk)error
}
