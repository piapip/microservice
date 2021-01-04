package handlers

import (
	"io"
	"net/http"
	"path/filepath"
	"strconv"

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

// UploadREST implements the http.Handler interface
func (f *Files) UploadREST(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	filename := vars["filename"]

	f.logger.Info("Handle POST", "id", id, "filename", filename)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters

	// check that the filepath is correct
	if id == "" || filename == "" {
		f.invalidURI(req.URL.String(), res)
		return
	}

	f.saveFile(id, filename, res, req.Body)
}

// UploadMultipart -
func (f *Files) UploadMultipart(res http.ResponseWriter, req *http.Request) {
	// 128*1024 is the amount of data will be held in the memory, the rest will be written on a temp file on the disk.
	// Reading from the memory is a lot faster than reading from files so need to make a decision of how much memory do I want
	// Like when you expect a 10GB worth of memory for a potential upload, you don't allocate 10GB to a request, too RAM-consuming.
	err := req.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.logger.Error("Bad request", "error", err)
		http.Error(res, "Expected multipart form request", http.StatusBadRequest)
		return
	}

	// After that we will have FormValue and FormFile available to us via req.
	// We will receive id and from the Multipart request
	id, idErr := strconv.Atoi(req.FormValue("id"))
	f.logger.Info("Process form for id", "id", id)

	if idErr != nil {
		f.logger.Error("Bad request", "error", err)
		http.Error(res, "Expected integer id", http.StatusBadRequest)
		return
	}

	file, multipartHeader, err := req.FormFile("file")
	if err != nil {
		f.logger.Error("Bad request", "error", err)
		http.Error(res, "Expected file", http.StatusBadRequest)
		return
	}

	f.saveFile(req.FormValue("id"), multipartHeader.Filename, res, file)
}

func (f *Files) invalidURI(url string, res http.ResponseWriter) {
	f.logger.Error("Invalid path", "path", url)
	http.Error(res, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

func (f *Files) saveFile(id, path string, res http.ResponseWriter, req io.ReadCloser) {
	f.logger.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, req)
	if err != nil {
		f.logger.Error("Unable to save file", "error", err)
		http.Error(res, "Unable to save file", http.StatusInternalServerError)
	}
}
