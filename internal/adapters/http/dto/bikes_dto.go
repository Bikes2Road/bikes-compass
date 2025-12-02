package dto

type BikesSearchRequest struct {
	Name  string `form:"name" binding:"required" example:"Suzuki Vstrom"`
	Page  int64  `form:"page" binding:"required" example:"1"`
	Cant  int64  `form:"cant" binding:"required" example:"10"`
	Brand string `form:"brand" binding:"required" example:"Suzuki"`
}

type BikesSearchResponse struct {
}

type BykeInfoRequest struct {
	HashByke string `uri:"hash_byke" binding:"required" example:"123abc456def"`
}

type BykeInfoResponse struct {
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

type PlaceHolderRequest struct {
	NameByke string `form:"name" validate:"required"`
}

type Photo struct {
	Url    string `json:"url" example:"http://photo_url.test"`
	Height int    `json:"height" example:"123"`
	Width  int    `json:"width" example:"123"`
	Key    string `json:"key" example:"/key/photo"`
}
