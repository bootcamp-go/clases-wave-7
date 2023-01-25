package response

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Ok(message string, data any) *Response {
	return &Response{Message: message, Data: data}
}

func Err(message string) *Response {
	return &Response{Message: message, Data: nil}
}