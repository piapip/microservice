package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/piapip/microservice/product-images/files"
)

// Files is a handler for reading and writing files
type Files struct {
	logger hclog.Logger
	store  files.Storage
}

// NewFiles creates a new Files handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, logger: l}
}

// ServeHTTP implements the http.Handler interface
func (f *Files) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	filename := vars["filename"]

	f.logger.Info("Handle POST", "id", id, "filename", filename)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters

	// check that the filepath is correct
	if id == "" || filename == "" {
		f.invalidURL(req.URL.String(), res)
		return
	}

	f.saveFile(id, filename, res, req)
}

func (f *Files) invalidURL(url string, res http.ResponseWriter) {
	f.logger.Error("Invalid path", "path", url)
	http.Error(res, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

func (f *Files) saveFile(id, path string, res http.ResponseWriter, req *http.Request) {
	f.logger.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, req.Body)
	if err != nil {
		f.logger.Error("Unable to save file", "error", err)
		http.Error(res, "Unable to save file", http.StatusInternalServerError)
	}
}
