package handlers

type HttpResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}
