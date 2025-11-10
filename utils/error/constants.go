package error

import (
	"net/http"
)

const SuccessStatus = false

const (
	ErrorBadRequest         = "error_bad_request"
	ErrorInvalidPage        = "error_invalid_page"
	ErrorInvalidCant        = "error_invalid_cant"
	ErrorInvalidStringBike  = "error_invalid_string_bike"
	ErrorBikesNotFound      = "error_bikes_not_found"
	ErrorBykeNotFound       = "error_byke_not_found"
	ErrorUpdateByke         = "error_update_byke"
	ErrorDeleteByke         = "error_delete_byke"
	ErrorUnauthorized       = "error_unauthorized"
	ErrorUnexpected         = "error_unexpected"
	ErrorMongoFindAll       = "error_mongo_find_all"
	ErrorMongoFind          = "error_mongo_find"
	ErrorR2Url              = "error_r2_generating_url"
	ErrorR2KeyEmpty         = "error_r2_key_empty"
	ErrorInvalidQueryParams = "error_query_params_invalids"
	ErrorInvalidPathParams  = "error_path_params_invalid"
	ErrorInvalidPathParam   = "error_path_param_invalid"
)

type ErrorInfo struct {
	Success bool
	Code    int
	Message string
	Error   error
}

var ErrorMappingsResponse = map[string]ErrorInfo{
	ErrorBadRequest: {
		Success: SuccessStatus,
		Code:    http.StatusBadRequest,
		Message: "Could Not interpret request",
	},
	ErrorInvalidPage: {
		Success: SuccessStatus,
		Code:    http.StatusBadRequest,
		Message: "Page could be greather than 1",
	},
	ErrorInvalidCant: {
		Success: SuccessStatus,
		Code:    http.StatusBadRequest,
		Message: "Cant could'nt greather than 30",
	},
	ErrorInvalidStringBike: {
		Success: SuccessStatus,
		Code:    http.StatusBadRequest,
		Message: "Byke name can only contain letters",
	},
	ErrorBikesNotFound: {
		Success: SuccessStatus,
		Code:    http.StatusNotFound,
		Message: "Bikes not found",
	},
	ErrorBykeNotFound: {
		Success: SuccessStatus,
		Code:    http.StatusNotFound,
		Message: "Byke %s not found",
	},
	ErrorUpdateByke: {
		Success: SuccessStatus,
		Code:    http.StatusBadRequest,
		Message: "Error updating byke %s",
	},
	ErrorDeleteByke: {
		Success: SuccessStatus,
		Code:    http.StatusBadRequest,
		Message: "Error deleting byke %s",
	},
	ErrorUnauthorized: {
		Success: SuccessStatus,
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	},
	ErrorInvalidQueryParams: {
		Success: SuccessStatus,
		Code:    http.StatusBadRequest,
		Message: "%s",
	},
	ErrorInvalidPathParams: {
		Success: SuccessStatus,
		Code:    http.StatusBadRequest,
		Message: "%s",
	},
	ErrorInvalidPathParam: {
		Success: SuccessStatus,
		Code:    http.StatusBadRequest,
		Message: "Path Param is not valid, only letters and numbers",
	},
	ErrorUnexpected: {
		Success: SuccessStatus,
		Code:    http.StatusInternalServerError,
		Message: "Unexpected error",
	},
}
