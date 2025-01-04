package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheusandrade23/go-bid/internal/store/pgstore"
)

type ProductService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewProductService(pool *pgxpool.Pool) ProductService {
	return ProductService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (ps *ProductService) CreateProduct(
	ctx context.Context, 
	sellerId uuid.UUID, 
	productName, 
	description string, 
	basePrice float64, 
	auctionEnd time.Time,
) (uuid.UUID, error) {

	id, err := ps.queries.CreateProduct(ctx, pgstore.CreateProductParams{
		SellerID: sellerId,
		ProductName: productName,
		BasePrice: basePrice,
		AuctionEnd: auctionEnd,
		Description: description,
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}