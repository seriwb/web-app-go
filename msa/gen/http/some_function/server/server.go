// Code generated by goa v3.0.4, DO NOT EDIT.
//
// some-function HTTP server
//
// Command:
// $ goa gen github.com/seriwb/front-bff-msa/msa/design

package server

import (
	"context"
	"net/http"

	somefunction "github.com/seriwb/front-bff-msa/msa/gen/some_function"
	goahttp "goa.design/goa/v3/http"
)

// Server lists the some-function service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the some-function service endpoints.
func New(
	e *somefunction.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"./gen/http/openapi.json", "GET", "/openapi.json"},
		},
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "some-function" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
}

// Mount configures the mux to serve the some-function endpoints.
func Mount(mux goahttp.Muxer) {
	MountGenHTTPOpenapiJSON(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/http/openapi.json")
	}))
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}