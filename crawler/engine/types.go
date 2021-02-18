package engine

type Request struct {
	Url       string
	ParseFunc func(string, ...interface{}) ParseResult
	Params    interface{}
}

type ParseResult struct {
	Request []Request
	Items   []interface{}
}
