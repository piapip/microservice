package data

import (
	"encoding/json"
	"io"
)

// ToJSON will convert list of given struct/interface to JSON format
func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	// All kind of funny shenanigan with i before encoding will go here. Like I'll only encode those goods that have a SKU.
	// Doing something with i here. I don't know how to though... Gotta look at it later...
	return encoder.Encode(i)
}

// FromJSON converts items in JSON format from the stream to the given struct/interface
func FromJSON(i interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	// This one is quite confusing tbh.
	return decoder.Decode(i)
}
