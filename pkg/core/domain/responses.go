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
