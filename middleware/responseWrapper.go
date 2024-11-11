package middleware

type ResponseWrapper struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}
