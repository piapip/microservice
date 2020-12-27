package handlers

import (
	"compress/gzip"
	"net/http"
)

// GzipHandler is a handler for zipper
type GzipHandler struct {
}

// WrappedResponseWriter - wrapper for zipper so won't expose it?
type WrappedResponseWriter struct {
	res http.ResponseWriter
	gw  *gzip.Writer
}

// NewWrappedResponseWriter - create new WrappedResponseWriter
func NewWrappedResponseWriter(res http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(res)

	return &WrappedResponseWriter{res: res, gw: gw}
}

// Header - return header of the zipper
func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.res.Header()
}

// Write - let the zipper write
func (wr *WrappedResponseWriter) Write(data []byte) (int, error) {
	return wr.res.Write(data)
}

// WriteHeader - ...
func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
	wr.res.WriteHeader(statuscode)
}

// Flush - ...
func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}

// GzipMiddleware will ...
func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return nil
}
