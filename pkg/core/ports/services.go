package ports

import (
	"context"

	"github.com/Bikes2Road/bikes-compass/pkg/core/domain"
)

type GetAllBikes interface {
	Execute(ctx context.Context, request domain.GetAllBikesRequest, path string) (*domain.GetAllResponseSuccess, *domain.ResponseHttpError)
}

type GetByke interface {
	Execute(ctx context.Context, requestByke domain.SearchBykeRequest, pathRequest string) (*domain.GetBykeResponseSuccess, *domain.ResponseHttpError)
}
