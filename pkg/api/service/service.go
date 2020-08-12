package service

import (
	"github.com/sid-sun/notes-api/pkg/api/contract/db"
	"github.com/sid-sun/notes-api/pkg/api/store"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
)

// Service Interface defines a service spec
type Service interface {
	Create(id string, data db.Data) error
	Exists(id string) bool
	Delete(id string)
	Get(id string) db.Data
}

// NotesService implements Service with store
type NotesService struct {
	logger *zap.Logger
	store  store.Store
}

// Create creates a new record in DB, handling translations
func (n NotesService) Create(id string, data db.Data) error {
	hash := sha3.Sum256([]byte(id))
	n.store.Put(string(hash[:]), data)
	return nil
}

// Get fetches the data corresponding to id in store
// and returns a db data, handling translations
func (n NotesService) Get(id string) db.Data {
	hash := sha3.Sum256([]byte(id))
	return n.store.Get(string(hash[:]))
}

// Delete deletes the data corresponding to id in store
// handing translations and deleting nothing
func (n NotesService) Delete(id string) {
	hash := sha3.Sum256([]byte(id))
	n.store.Delete(string(hash[:]))
}

// Exists gets data with id from DB, checks it against zero values
// and returns true if the record is non-zero
func (n NotesService) Exists(id string) bool {
	hash := sha3.Sum256([]byte(id))
	d := n.store.Get(string(hash[:]))
	// If empty, data does NOT exist so NOT it
	return !d.IsEmpty()
}

// NewNotesService creates a new instance of NotesService
func NewNotesService(st store.Store, lg *zap.Logger, ) NotesService {
	return NotesService{
		logger: lg,
		store:  st,
	}
}
