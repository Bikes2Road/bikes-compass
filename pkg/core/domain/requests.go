package domain

type GetAllBikesRequest struct {
	Name  string `form:"name" validate:"required"`
	Page  int64  `form:"page" validate:"required"`
	Cant  int64  `form:"cant" validate:"required"`
	Brand string `form:"brand" validate:"required"`
}

type SearchBykeRequest struct {
	HashByke string `uri:"hash_byke" binding:"required"`
}

type PlaceHolderRequest struct {
	NameByke string `form:"name" validate:"required"`
}
