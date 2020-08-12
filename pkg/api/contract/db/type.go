package db

import "bytes"

// Data defines structure of db data
type Data struct {
	AAD  []byte   `json:"aad" bson:"aad"`
	Hash [32]byte `json:"hash" bson:"hash"`
	Note []byte   `json:"note" bson:"note"`
}

// IsEmpty checks if data is empty
func (d Data) IsEmpty() bool {
	return bytes.Equal(d.Note, []byte{})
}

// NewDataInstance creates a new data instance, initialising fields to default values
func NewDataInstance(aad []byte, hash [32]byte, note []byte) Data {
	return Data{
		AAD:  aad,
		Hash: hash,
		Note: note,
	}
}
