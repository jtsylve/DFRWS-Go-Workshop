package recycle

import (
	"errors"
	"io"
	"time"
)

// ErrInvalid is returned when the data passed to a Decode function
// is not properly formatted.
var ErrInvalid = errors.New("Invalid recycle bin file.")

// Metadata gives information about a deleted file
type Metadata struct {
	Name    string    // Original File Name
	Size    int64     // Original File Size
	Deleted time.Time // The Date And Time The File Was Deleted
}

// DecodeI takes a Windows Vista+ $RECYCLE.BIN $I file as r and
// returns associated Metadata for the deleted file. ErrInvalid
// is returned if r is not a valid $I file.
func DecodeI(r io.Reader) (Metadata, error) {
	// TODO: Implement me!
	return Metadata, nil
}
