package repository

import "context"

type Database interface {
	GetMerchantIdByUserIdChatat(ctx context.Context, merchantId string) string
}
