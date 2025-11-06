package domain

type GetAllBikesRequest struct {
	Name string `form:"name" validate:"required"`
	Page int64  `form:"page" validate:"required"`
	Cant int64  `form:"cant" validate:"required"`
}

type SearchBykeRequest struct {
	HashByke string `json:"hash_byke" validate:"required"`
}
