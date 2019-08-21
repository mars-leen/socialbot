package pinmodels

type ResponseError struct {
	APIErrorCode int `json:"api_error_code"`
	Message string `json:"message"`
}

type BaseResponse struct {
	ResourceResponse struct {
		ResponseError *ResponseError `json:"error"`
	} `json:"resource_response"`
}

