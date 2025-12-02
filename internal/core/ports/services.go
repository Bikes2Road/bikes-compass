package ports

import (
	"context"

	"github.com/Bikes2Road/bikes-compass/internal/core/domain"
)

type GetAllBikes interface {
	Execute(ctx context.Context, request domain.GetAllBikesRequest, path string) (*domain.GetAllResponseSuccess, *domain.ResponseHttpError)
}

type GetByke interface {
	Execute(ctx context.Context, requestByke domain.SearchBykeRequest, pathRequest string) (*domain.GetBykeResponseSuccess, *domain.ResponseHttpError)
}

type PlaceHolder interface {
	Execute(ctx context.Context, requestPlaceHolder domain.PlaceHolderRequest) (*domain.PlaceHolderResponseSuccess, *domain.ResponseHttpError)
}
