package api

type ApiError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func CreateApiError(message string, code int) ApiError {
	return ApiError{
		Message: message,
		Code:    code,
	}
}
