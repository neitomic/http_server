package response

import "net/http"

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HttpResp struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Error      *APIError   `json:"error,omitempty"`
	HttpCode   int         `json:"-"`
	HttpHeader http.Header `json:"-"`
}

func SuccessResp(data interface{}) HttpResp {
	return HttpResp{
		Success:    true,
		Data:       data,
		Error:      nil,
		HttpCode:   http.StatusOK,
		HttpHeader: nil,
	}
}

func ErrorResp(httpCode int, error *APIError) HttpResp {
	return HttpResp{
		Success:    false,
		Data:       nil,
		Error:      error,
		HttpCode:   httpCode,
		HttpHeader: nil,
	}
}
