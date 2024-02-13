package dto

type HttpMsg struct {
	Error *HttpError `json:"error,omitempty"`
	Data  any        `json:"data,omitempty"`
}

type HttpError struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewHttpErrorMsg(status bool, code int, msg string) *HttpMsg {
	return &HttpMsg{
		Error: &HttpError{status, code, msg},
	}
}

func NewHttpDataMsg(data any) *HttpMsg {
	return &HttpMsg{
		Data: data,
	}
}
