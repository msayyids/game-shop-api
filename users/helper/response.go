package helper

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omiempty"`
}

func NewResponse(status int, detail string, data interface{}) Response {
	return Response{StatusCode: status, Message: detail, Data: data}
}
