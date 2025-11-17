package domain

// swagger:model GetAllResponseSuccess
// GetAllResponseSuccess representa la respuesta exitosa al listar motos.
type GetAllResponseSuccess struct {
	// Indica si la petición fue exitosa
	Success bool `json:"success" validate:"required" example:"true"`
	// Lista de motos encontradas
	Data []*BykeReponse `json:"data" validate:"required" swaggertype:"array,object"`
	// Número total de registros encontrados
	Total int64 `json:"total" validate:"required" example:"10"`
}

type GetBykeResponseSuccess struct {
	// Indica si la petición fue exitosa
	Success bool `json:"success" validate:"required" example:"true"`
	// Lista de motos encontradas
	Data *FullBykeResponse `json:"data" validate:"required" swaggertype:"array,object"`
	// Número total de registros encontrados
	Total int64 `json:"total" validate:"required" example:"10"`
}

type PlaceHolderResponseSuccess struct {
	Success bool     `json:"success" validate:"required" example:"true"`
	Data    []string `json:"data" validate:"required" example:"[\"Ducati\",\"BMW\"]"`
	Total   int64    `json:"total" validate:"required" example:"2"`
}

// swagger:model ResponseHttpError
// ResponseHttpError representa la estructura de un error HTTP estándar.
type ResponseHttpError struct {
	// Código de error HTTP
	Code int `json:"code" validate:"required" example:"400"`
	// Identificador del tipo de error
	Error string `json:"error" validate:"required" example:"bad_request"`
	// Indica si la operación fue exitosa (en este caso false)
	Success bool `json:"success" validate:"required" example:"false"`
	// Mensaje descriptivo del error
	Message string `json:"message" validate:"required" example:"error with request"`
}

type FullBykeResponse struct {
	Ref           string    `json:"ref" bson:"ref" example:"1234"`
	HashByke      string    `json:"hash_byke" bson:"hash_byke" example:"abcd1234"`
	FullName      string    `json:"full_name" bson:"full_name" example:"Yamaha MT-03"`
	Brand         string    `json:"brand" bson:"brand"`
	Model         string    `json:"model" bson:"model"`
	Cylinder      string    `json:"cylinder" bson:"cylinder"`
	Engine        string    `json:"engine" bson:"engine"`
	HorsePower    string    `json:"horse_power" bson:"horse_power"`
	Weight        string    `json:"weight" bson:"weight"`
	CityRegister  string    `json:"city_register" bson:"city_register"`
	Extras        []string  `json:"extras" bson:"extras,omitempty"`
	DateFound     int       `json:"date_found" bson:"date_found"`
	DateSoat      string    `json:"date_soat" bson:"date_soat"`
	DateTecnico   string    `json:"date_tecnico" bson:"date_tecnico"`
	PageInstagram string    `json:"page_instagram" bson:"page_instagram"`
	UrlPost       string    `json:"url_post" bson:"url_post"`
	YearModel     int       `json:"year_model" bson:"year_model" example:"2020"`
	Kilometers    int       `json:"km" bson:"km" example:"1235"`
	Price         int       `json:"price" bson:"price" example:"25000000"`
	Location      string    `json:"location" bson:"location" example:"Bogotá D.C"`
	DatePublish   int       `json:"date_publish" bson:"date_publish" example:"1731081212"`
	Photos        [][]Photo `json:"photos" bson:"photos" swaggertype:"array,array,object"`
	Torque        string    `json:"torque" bson:"torque"`
}

// swagger:model BykeReponse
// BykeReponse representa la información detallada de una moto.
type BykeReponse struct {
	// Identificador de referencia de la moto
	Ref string `json:"ref" bson:"ref" example:"1234"`
	// Hash único de la moto
	HashByke string `json:"hash_byke" bson:"hash_byke" example:"abcd1234"`
	// Nombre completo del modelo
	FullName string `json:"full_name" bson:"full_name" example:"Yamaha MT-03"`
	// Año del modelo
	YearModel int `json:"year_model" bson:"year_model" example:"2020"`
	// Kilometraje de la moto
	Kilometers int `json:"km" bson:"km" example:"1235"`
	// Precio de la moto
	Price int `json:"price" bson:"price" example:"25000000"`
	// Ciudad o región donde está ubicada
	Location string `json:"location" bson:"location" example:"Bogotá D.C"`
	// Fecha de publicación (timestamp)
	DatePublish int `json:"date_publish" bson:"date_publish" example:"1731081212"`
	// Fotos asociadas a la moto
	Photos [][]Photo `json:"photos" bson:"photos" swaggertype:"array,array,object"`
}

type BykeName struct {
	FullName string `json:"full_name" bson:"full_name"`
}
