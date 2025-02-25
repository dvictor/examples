// Code generated by goa v3.0.2, DO NOT EDIT.
//
// storage HTTP server
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package server

import (
	"context"
	"mime/multipart"
	"net/http"

	storage "goa.design/examples/cellar/gen/storage"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the storage service endpoint HTTP handlers.
type Server struct {
	Mounts      []*MountPoint
	List        http.Handler
	Show        http.Handler
	Add         http.Handler
	Remove      http.Handler
	Rate        http.Handler
	MultiAdd    http.Handler
	MultiUpdate http.Handler
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

// StorageMultiAddDecoderFunc is the type to decode multipart request for the
// "storage" service "multi_add" endpoint.
type StorageMultiAddDecoderFunc func(*multipart.Reader, *[]*storage.Bottle) error

// StorageMultiUpdateDecoderFunc is the type to decode multipart request for
// the "storage" service "multi_update" endpoint.
type StorageMultiUpdateDecoderFunc func(*multipart.Reader, **storage.MultiUpdatePayload) error

// New instantiates HTTP handlers for all the storage service endpoints.
func New(
	e *storage.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
	storageMultiAddDecoderFn StorageMultiAddDecoderFunc,
	storageMultiUpdateDecoderFn StorageMultiUpdateDecoderFunc,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"List", "GET", "/storage"},
			{"Show", "GET", "/storage/{id}"},
			{"Add", "POST", "/storage"},
			{"Remove", "DELETE", "/storage/{id}"},
			{"Rate", "POST", "/storage/rate"},
			{"MultiAdd", "POST", "/storage/multi_add"},
			{"MultiUpdate", "PUT", "/storage/multi_update"},
		},
		List:        NewListHandler(e.List, mux, dec, enc, eh),
		Show:        NewShowHandler(e.Show, mux, dec, enc, eh),
		Add:         NewAddHandler(e.Add, mux, dec, enc, eh),
		Remove:      NewRemoveHandler(e.Remove, mux, dec, enc, eh),
		Rate:        NewRateHandler(e.Rate, mux, dec, enc, eh),
		MultiAdd:    NewMultiAddHandler(e.MultiAdd, mux, NewStorageMultiAddDecoder(mux, storageMultiAddDecoderFn), enc, eh),
		MultiUpdate: NewMultiUpdateHandler(e.MultiUpdate, mux, NewStorageMultiUpdateDecoder(mux, storageMultiUpdateDecoderFn), enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "storage" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.List = m(s.List)
	s.Show = m(s.Show)
	s.Add = m(s.Add)
	s.Remove = m(s.Remove)
	s.Rate = m(s.Rate)
	s.MultiAdd = m(s.MultiAdd)
	s.MultiUpdate = m(s.MultiUpdate)
}

// Mount configures the mux to serve the storage endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListHandler(mux, h.List)
	MountShowHandler(mux, h.Show)
	MountAddHandler(mux, h.Add)
	MountRemoveHandler(mux, h.Remove)
	MountRateHandler(mux, h.Rate)
	MountMultiAddHandler(mux, h.MultiAdd)
	MountMultiUpdateHandler(mux, h.MultiUpdate)
}

// MountListHandler configures the mux to serve the "storage" service "list"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "list" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeListResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountShowHandler configures the mux to serve the "storage" service "show"
// endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/{id}", f)
}

// NewShowHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "show" endpoint.
func NewShowHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeShowRequest(mux, dec)
		encodeResponse = EncodeShowResponse(enc)
		encodeError    = EncodeShowError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAddHandler configures the mux to serve the "storage" service "add"
// endpoint.
func MountAddHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage", f)
}

// NewAddHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "add" endpoint.
func NewAddHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAddRequest(mux, dec)
		encodeResponse = EncodeAddResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountRemoveHandler configures the mux to serve the "storage" service
// "remove" endpoint.
func MountRemoveHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/storage/{id}", f)
}

// NewRemoveHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "remove" endpoint.
func NewRemoveHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveRequest(mux, dec)
		encodeResponse = EncodeRemoveResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "remove")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountRateHandler configures the mux to serve the "storage" service "rate"
// endpoint.
func MountRateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/rate", f)
}

// NewRateHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "rate" endpoint.
func NewRateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRateRequest(mux, dec)
		encodeResponse = EncodeRateResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "rate")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountMultiAddHandler configures the mux to serve the "storage" service
// "multi_add" endpoint.
func MountMultiAddHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/multi_add", f)
}

// NewMultiAddHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "multi_add" endpoint.
func NewMultiAddHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeMultiAddRequest(mux, dec)
		encodeResponse = EncodeMultiAddResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "multi_add")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountMultiUpdateHandler configures the mux to serve the "storage" service
// "multi_update" endpoint.
func MountMultiUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/storage/multi_update", f)
}

// NewMultiUpdateHandler creates a HTTP handler which loads the HTTP request
// and calls the "storage" service "multi_update" endpoint.
func NewMultiUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeMultiUpdateRequest(mux, dec)
		encodeResponse = EncodeMultiUpdateResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "multi_update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
