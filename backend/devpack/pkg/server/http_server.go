package pkgserver

type HttpServer interface {
	GET(path string, handler Handler)
	POST(path string, handler Handler)
	PUT(path string, handler Handler)
	DELETE(path string, handler Handler)
	PATCH(path string, handler Handler)
	Start(port int) error
	Stop() error
}
