package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/piapip/microservice/product-images/files"
)

// Files is a handler for reading and writing files
type Files struct {
	logger hclog.Logger
	store  files.Storage // haven't implemented yet lolz
}

// NewFiles creates a new Files handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, logger: l}
}
