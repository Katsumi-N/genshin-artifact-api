package domain

type HttpClient interface {
	SendRequest(method string, baseURL string, endpointPath string, data string, query map[string]string) (body []byte, statusCode int, err error)
}
