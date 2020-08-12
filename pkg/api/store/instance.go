package store

import (
	"github.com/sid-sun/notes-api/pkg/api/contract/db"
	"go.uber.org/zap"
)

// InstanceInterface defines a db instance interface
type InstanceInterface interface {
	Get(string) db.Data
	Delete(string)
	Put(string, db.Data)
}

// NewInstance creates a new instance for db
func NewInstance(logger *zap.Logger) InstanceInterface {
	return Instance{logger: logger, data: make(map[string]db.Data)}
}

// Instance implements InstanceInterface with map
type Instance struct {
	data   map[string]db.Data
	logger *zap.Logger
}

// Delete deletes the db Data instance corresponding to id
func (i Instance) Delete(id string) {
	i.data[id] = db.Data{}
}

// Get returns a db Data instance corresponding to id
func (i Instance) Get(id string) db.Data {
	return i.data[id]
}

// Put unconditionally sets db record of id to provided data
func (i Instance) Put(id string, data db.Data) {
	i.data[id] = data
}
