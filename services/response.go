package services

type HttpResponse struct {
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data,omitempty"`
	Error      string `json:"error,omitempty"`
	Errors     any    `json:"errors,omitempty"`
}
