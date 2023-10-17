package helper

type ErrorResponse struct {
	StatusCode int
	Detail     string
}

func NewErrorResponse(status int, detail string) ErrorResponse {
	return ErrorResponse{StatusCode: status, Detail: detail}
}
