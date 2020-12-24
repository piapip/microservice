package files

import "path/filepath"

// Local is an implementation of the Storage interface which works with the
// local disk on the current computer
type Local struct {
	maxFileSize int // maximum numbber of bytes for files
	basePath    string
}

// NewLocal creates a new Local filesytem with the given base path
// basePath is the base directory to save files to
// maxSize is the max number of bytes that a file can be
func NewLocal(basePath string, maxSize int) (*Local, error) {
	path, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	// return &Local{basePath: p}, nil
	return &Local{basePath: path, maxFileSize: maxSize}, nil
}

func (l *Local) fullPath(path string) string {
	// append the given path to the base path
	return filepath.Join(l.basePath, path)
}
