package exception

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NotFoundError(message string) *AppError {
	return &AppError{
		Code:    404,
		Message: message,
	}
}
