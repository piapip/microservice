package files

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupLocal(t *testing.T) (*Local, string, func()) {
	// create a temporary directory
	dir, err := ioutil.TempDir("", "files")
	if err != nil {
		t.Fatal(err)
	}

	local, err := NewLocal(dir, 100000)
	if err != nil {
		t.Fatal(err)
	}

	return local, dir, func() {
		// cleanup function
		// os.RemoveAll(dir) // why do we have to call RemoveAll? Doesn't TempDir will remove itself when the test is over?
	}
}

func TestSavesContentsOfReader(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "Hello World"

	// setup testing environment
	local, dir, cleanup := setupLocal(t)
	defer cleanup()

	err := local.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	// check if the file is correctly written
	f, err := os.Open(filepath.Join(dir, savePath))
	assert.NoError(t, err)

	// check the content of the file
	d, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	assert.Equal(t, fileContents, string(d))
}

func TestGetsContentsAndWritesToWriter(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "Hello World"

	// setup testing environment
	local, _, cleanup := setupLocal(t)
	defer cleanup()

	// save a file
	err := local.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	// Read that file
	r, err := local.Get(savePath)
	assert.NoError(t, err)
	defer r.Close()

	// read the full contents of the reader
	d, err := ioutil.ReadAll(r)
	assert.Equal(t, fileContents, string(d))
}
