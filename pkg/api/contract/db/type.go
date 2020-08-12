package db

import (
	"bytes"
	"go.mongodb.org/mongo-driver/bson"
)

type Query struct {
	ID   string   `json:"id" bson:"id"`
}

// ToBSON converts data to BSON representation
func (q Query) ToBSON() ([]byte, error) {
	return bson.Marshal(q)
}


// Data defines structure of db data
type Data struct {
	ID   string   `json:"id" bson:"id"`
	AAD  []byte   `json:"aad" bson:"aad"`
	Hash [32]byte `json:"hash" bson:"hash"`
	Note []byte   `json:"note" bson:"note"`
}

// ToBSON converts data to BSON representation
func (d Data) ToBSON() ([]byte, error) {
	return bson.Marshal(d)
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
