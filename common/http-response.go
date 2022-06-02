package common

type HttpResponse struct {
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data,omitempty"`
	Message    string `json:"message,omitempty"`
	Error      any    `json:"error,omitempty"`
}
