package response

type Response struct {
	Message string
	Data    interface{}
}

func Ok(message string, data interface{}) *Response {
	return &Response{Message: message, Data: data}
}

func Err(err error) *Response {
	return &Response{Message: err.Error(), Data: nil}
}