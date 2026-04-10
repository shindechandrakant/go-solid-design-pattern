package design_pattern

import "fmt"

type HttpRequest struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    string
}

type HttpRequestBuilder struct {
	request *HttpRequest
}

func NewHttpRequestBuilder() *HttpRequestBuilder {
	return &HttpRequestBuilder{
		request: &HttpRequest{
			Headers: make(map[string]string),
		},
	}
}

func (b *HttpRequestBuilder) URL(url string) *HttpRequestBuilder {
	b.request.URL = url
	return b
}

func (b *HttpRequestBuilder) Method(method string) *HttpRequestBuilder {
	b.request.Method = method
	return b
}

func (b *HttpRequestBuilder) Body(body string) *HttpRequestBuilder {
	b.request.Body = body
	return b
}

func (b *HttpRequestBuilder) Header(key, value string) *HttpRequestBuilder {
	b.request.Headers[key] = value
	return b
}

func (b *HttpRequestBuilder) Build() (*HttpRequest, error) {
	if b.request.URL == "" {
		return nil, fmt.Errorf("URL is required")
	}
	return b.request, nil
}

func main() {

	//body := struct {
	//	Name string
	//}{
	//	Name: "Chandrakant",
	//}
	//
	//request, _ :=
	//	NewHttpRequestBuilder()
	//.URL("Https://chandrakant.dev")
	//.Method("POST")
	//.Header("Content-Type", "application/json")
	//.Body(body.Name)
	//.Build()
}
