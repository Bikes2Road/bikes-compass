package domain

type GetAllResponseSuccess struct {
	Status string         `json:"status" validate:"required"`
	Data   []*BykeReponse `json:"data" validate:"required"`
	Total  int64          `json:"total" validate:"required"`
}

type GetAllResponseError struct {
	Status  string `json:"status" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type BykeReponse struct {
	Ref         string    `json:"ref"`
	HashByke    string    `json:"hash_byke"`
	FullName    string    `json:"full_name"`
	YearModel   int       `json:"year_model"`
	Kilometers  int       `json:"km"`
	Price       int       `json:"price"`
	Location    string    `json:"location"`
	DatePublish int       `json:"date_publish"`
	Photos      [][]Photo `json:"photos"`
}
