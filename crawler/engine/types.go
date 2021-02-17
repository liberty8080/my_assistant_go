package engine

type Request struct {
	Url       string
	ParseFunc func(string) ParseResult
}

type ParseResult struct {
	Request []Request
	Items   []interface{}
}
