package httpclient

type response struct {
	contentLength int
	body          []byte
	statusCode    int
	header        map[string][]string
}

type Response interface {
	ContentLength() int
	StatusCode() int
	Body() []byte
	Header() map[string][]string
}

func (r *response) ContentLength() int {
	return r.contentLength
}

func (r *response) StatusCode() int {
	return r.statusCode
}

func (r *response) Body() []byte {
	return r.body
}

func (r *response) Header() map[string][]string {
	return r.header
}
