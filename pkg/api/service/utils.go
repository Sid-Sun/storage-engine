package service

import (
	"bytes"
	"github.com/sid-sun/notes-api/pkg/api/contract/db"
)

// DataIsEmpty checks if the provided data instance is empty or not
func DataIsEmpty(data db.Data) bool {
	return bytes.Equal(data.Note, []byte{})
}
