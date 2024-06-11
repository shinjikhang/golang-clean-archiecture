package common

type Response struct {
	//Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination,omitempty"` // neu key la nil thi se khong hien thi trong json
	Filter     interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *Response {
	return &Response{
		Data:       data,
		Pagination: paging,
		Filter:     filter,
	}
}

func SimpleSuccessResponse(data interface{}) *Response {
	return &Response{Data: data}
}

func ErrorResponse(code int, err error) *Response {
	return &Response{
		//Code: code,
		Data: struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		},
	}
}
