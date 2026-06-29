package pkgserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
)

type HttpAdapter struct {
	mux  *http.ServeMux
	port int
}

func NewHttpAdapter() *HttpAdapter {
	return &HttpAdapter{mux: http.NewServeMux()}
}

func (a *HttpAdapter) Handler() http.Handler {
	return a.mux
}

func (a *HttpAdapter) Start(port int) error {
	a.port = port
	fmt.Println("Starting server on port", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), a.mux)
}

func (a *HttpAdapter) Stop() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.port),
		Handler: a.mux,
	}
	return server.Shutdown(context.Background())
}

func (a *HttpAdapter) GET(path string, handler Handler) {
	a.register(http.MethodGet, path, handler)
}

func (a *HttpAdapter) POST(path string, handler Handler) {
	a.register(http.MethodPost, path, handler)
}

func (a *HttpAdapter) PUT(path string, handler Handler) {
	a.register(http.MethodPut, path, handler)
}

func (a *HttpAdapter) DELETE(path string, handler Handler) {
	a.register(http.MethodDelete, path, handler)
}

func (a *HttpAdapter) PATCH(path string, handler Handler) {
	a.register(http.MethodPatch, path, handler)
}

func (a *HttpAdapter) register(method, path string, handler Handler) {
	paramNames := extractParamNames(path)
	pattern := method + " " + toServeMuxPattern(path)
	a.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		serveHandler(w, r, handler, paramNames)
	})
}

func serveHandler(w http.ResponseWriter, r *http.Request, handler Handler, paramNames []string) {
	input, err := buildHandlerInput(r, paramNames)
	if err != nil {
		writeHandlerError(w, err)
		return
	}
	response, err := handler(input)
	if err != nil {
		writeHandlerError(w, err)
		return
	}
	writeJSON(w, response.StatusCode, response.Body)
}

func buildHandlerInput(r *http.Request, paramNames []string) (*HandlerInput, error) {
	params := make(map[string]string, len(paramNames))
	for _, name := range paramNames {
		params[name] = r.PathValue(name)
	}
	query := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			query[key] = values[0]
		}
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return &HandlerInput{
		Params: params,
		Query:  query,
		Body:   body,
	}, nil
}

func writeHandlerError(w http.ResponseWriter, err error) {
	var customError pkgerror.ErrorC
	if errors.As(err, &customError) {
		errorResponse := map[string]string{
			"message": err.Error(),
		}
		switch customError.Code() {
		case pkgerror.DomainErrorCode:
			writeJSON(w, http.StatusBadRequest, errorResponse)
		case pkgerror.NotFoundErrorCode:
			writeJSON(w, http.StatusNotFound, errorResponse)
		default:
			writeJSON(w, http.StatusInternalServerError, errorResponse)
		}
		return
	}
	writeJSON(w, http.StatusInternalServerError, map[string]string{
		"message": "Unexpected error",
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func extractParamNames(path string) []string {
	var names []string
	for _, part := range strings.Split(path, "/") {
		if strings.HasPrefix(part, ":") {
			names = append(names, part[1:])
		}
	}
	return names
}

func toServeMuxPattern(path string) string {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			parts[i] = "{" + part[1:] + "}"
		}
	}
	return strings.Join(parts, "/")
}
