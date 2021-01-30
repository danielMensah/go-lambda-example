package aws

type fn func(req Request) Response

func Handle(req Request, handler fn) Response {
	return handler(req)
}