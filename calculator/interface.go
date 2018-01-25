package calculator

type Calculator interface {
	Sum(*Request, *Response) error
}
