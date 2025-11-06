package domain

type Bike struct {
	ID            interface{} `json:"-" bson:"_id,omitempty"`
	HashByke      string      `json:"hash_byke" bson:"hash_byke"`
	Ref           string      `json:"ref" bson:"ref"`
	Brand         string      `json:"brand" bson:"brand"`
	Model         string      `json:"model" bson:"model"`
	FullName      string      `json:"full_name" bson:"full_name"`
	YearModel     int         `json:"year_model" bson:"year_model"`
	Cylinder      int         `json:"cylinder" bson:"cylinder"`
	Engine        string      `json:"engine" bson:"engine"`
	HorsePower    string      `json:"horse_power" bson:"horse_power"`
	Kilometers    int         `json:"km" bson:"km"`
	Weight        string      `json:"weight" bson:"weight"`
	CityRegister  string      `json:"city_register" bson:"city_register"`
	Extras        []string    `json:"extras" bson:"extras,omitempty"`
	DateFound     int         `json:"date_found" bson:"date_found"`
	DatePublish   int         `json:"date_publish" bson:"date_publish"`
	DateSoat      string      `json:"date_soat" bson:"date_soat"`
	DateTecnico   string      `json:"date_tecnico" bson:"date_tecnico"`
	Description   string      `json:"description" bson:"description,omitempty"`
	PageInstagram string      `json:"page_instagram" bson:"page_instagram"`
	Photos        [][]Photo   `json:"photos" bson:"photos"`
	UrlPost       string      `json:"url_post" bson:"url_post"`
	Price         int         `json:"price" bson:"price"`
	Location      string      `json:"location" bson:"location"`
}

type Photo struct {
	Url    string
	Height int
	Width  int
	Key    string
}
