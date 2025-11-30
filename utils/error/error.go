package error

import (
	"fmt"

	"github.com/Bikes2Road/bikes-compass/internal/core/domain"
)

type WrapperError struct {
	Type    string
	Message error
}

func MapError(typeError string, err error) *WrapperError {
	fmt.Printf("%v: %v\n", typeError, err)
	return &WrapperError{
		Type:    typeError,
		Message: err,
	}
}

func MapErrorResponse(typeError string, err error) *domain.ResponseHttpError {
	errorInfo, exists := ErrorMappingsResponse[typeError]
	if !exists {
		errorInfo = ErrorMappingsResponse[ErrorUnexpected]

		//fmt.Printf("%v: %v", ErrorUnexpected, err)
		return &domain.ResponseHttpError{
			Code:    errorInfo.Code,
			Error:   ErrorUnexpected,
			Success: errorInfo.Success,
			Message: errorInfo.Message,
		}
	}

	if typeError == ErrorInvalidQueryParams {
		return &domain.ResponseHttpError{
			Code:    errorInfo.Code,
			Error:   typeError,
			Success: errorInfo.Success,
			Message: fmt.Sprintf(errorInfo.Message, err),
		}
	}

	if typeError == ErrorInvalidPathParams {
		return &domain.ResponseHttpError{
			Code:    errorInfo.Code,
			Error:   typeError,
			Success: errorInfo.Success,
			Message: fmt.Sprintf(errorInfo.Message, err),
		}
	}

	//fmt.Printf("%v: %v", typeError, err)

	return &domain.ResponseHttpError{
		Code:    errorInfo.Code,
		Error:   typeError,
		Success: errorInfo.Success,
		Message: errorInfo.Message,
	}
}
