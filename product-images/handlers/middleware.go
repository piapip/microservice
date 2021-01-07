package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// GzipHandler is a handler for zipper
type GzipHandler struct {
}

// WrappedResponseWriter - we will treat it as another ResponseWriter so we will have to implement the interface http.ResponseWriter's methods later below
type WrappedResponseWriter struct {
	res http.ResponseWriter
	gw  *gzip.Writer
}

// NewWrappedResponseWriter - create new WrappedResponseWriter
func NewWrappedResponseWriter(res http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(res)

	return &WrappedResponseWriter{res: res, gw: gw}
}

// Header - return header of the zipper - implementing interface's method
func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.res.Header()
}

// Write - let the zipper write - implementing interface's method
func (wr *WrappedResponseWriter) Write(data []byte) (int, error) {
	return wr.res.Write(data)
}

// WriteHeader - implementing interface's method
func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
	wr.res.WriteHeader(statuscode)
}

// Flush - flush anything that hasn't been sent out on the steram
func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}

// GzipMiddleware will ...
func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// when an HTTP client REQUEST that it can handle content in a gzip,
		// it's going to send the header accept-encoding.
		// Definition:
		// Accept-encoding request HTTP header advertiser which content encoding, usually a compression algorithm, the client is able to understand.
		// So when we don't have this Accept-encoding or we have something which we don't understand, then we will send plain text back

		// get header information from Request
		// Here's a sample of Accept-Encoding:
		// Accept-Encoding: deflate, gzip;q=1.0, *, q=0.5
		// For simplicity, we won't separate each part of the header, we will compare string instead.
		if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			// create a gzip response
			wrapRes := NewWrappedResponseWriter(res)
			wrapRes.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrapRes, req)
			defer wrapRes.Flush()

			return
		}
		// if the thing we do doesn't relate to gzip then move on normally
		next.ServeHTTP(res, req)
	})
}
