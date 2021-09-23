package go_flat_http

type Response struct {
	err          error
	data         interface{}
	responseCode int
}

func (r *Response) ResponseCode() int {
	return r.responseCode
}

func (r Response) WithResponseCode(responseCode int) Response {
	r.responseCode = responseCode

	return r
}

func (r *Response) Data() interface{} {
	return r.data
}

func (r Response) WithData(data interface{}) Response {
	r.data = data

	return r
}

func (r *Response) Err() error {
	return r.err
}

func (r *Response) HasErr() bool {
	return r.err != nil
}

func (r Response) WithErr(err error) Response {
	r.err = err

	return r
}
