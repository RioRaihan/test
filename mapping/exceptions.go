package mapping

import (
	"net/http"
	"os"
)

//type MappingType struct {
//}

func Mapping(err error) {
	MapToHttpStatusCode(err)
	MapToErrorMessage(err)
}

func MapToHttpStatusCode(err error) int {
	var statusCode = http.StatusInternalServerError
	switch err {
	case ErrSystem:
		statusCode = http.StatusInternalServerError
	case ErrNotFound:
		statusCode = http.StatusNotFound
	case ErrBadRequest:
		statusCode = http.StatusBadRequest
	case ErrInvalidPayload:
		statusCode = http.StatusBadRequest
	case ErrEligibility:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}

func MapToErrorMessage(err error) string {
	var errorMessage = os.Getenv("ERRSYSTEM_TEXT")

	switch err {
	case ErrSystem:
		errorMessage = os.Getenv("ERRSYSTEM_TEXT")
	case ErrNotFound:
		errorMessage = os.Getenv("ERRNOTFOUND_TEXT")
	case ErrBadRequest:
		errorMessage = os.Getenv("ERRNOTFOUND_TEXT")
	case ErrEligibility:
		errorMessage = os.Getenv("ERRELIGIBILITY_TEXT")
	case ErrInvalidPayload:
		errorMessage = os.Getenv("ERRPAYLOAD_TEXT")
	default:
		errorMessage = os.Getenv("ERRSYSTEM_TEXT")
	}

	return errorMessage
}
