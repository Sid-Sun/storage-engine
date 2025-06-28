package service

import (
	"encoding/base64"
	"fmt"
	"github.com/sid-sun/storage-engine/pkg/api/contract/db"
	"github.com/sid-sun/storage-engine/pkg/api/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
)

// Service Interface defines a service spec
type Service interface {
	Create(id string, data db.Data) error
	Update(id string, data db.Data) error
	Exists(id string) (bool, error)
	Delete(id string) error
	Get(id string) (db.Data, error)
}

// NotesService implements Service with store
type NotesService struct {
	logger *zap.Logger
	store  store.Store
}

// Create creates a new record in DB, handling translations
func (n NotesService) Create(id string, data db.Data) error {
	hash := sha3.Sum256([]byte(id))
	data.ID = base64.StdEncoding.EncodeToString(hash[:])
	doc, err := data.ToBSON()
	if err != nil {
		n.logger.Sugar().Errorf("%s : %v", "[Service] [Create] [ToBSON]", err)
		return err
	}

	_, err = n.store.Insert(doc)
	if err != nil {
		return err
	}
	return nil
}

// Get fetches the data corresponding to id in store
// and returns a db data, handling translations
func (n NotesService) Get(id string) (db.Data, error) {
	hash := sha3.Sum256([]byte(id))
	q := db.Query{ID: base64.StdEncoding.EncodeToString(hash[:])}
	doc, err := q.ToBSON()
	if err != nil {
		n.logger.Sugar().Errorf("%s : %v", "[Service] [Get] [ToBSON]", err)
		return db.Data{}, err
	}

	res, err := n.store.Find(doc)
	if err != nil {
		return db.Data{}, err
	}
	var d db.Data
	err = res.Decode(&d)
	if err != nil && err != mongo.ErrNoDocuments {
		n.logger.Sugar().Errorf("%s : %v", "[Service] [Get] [Decode]", err)
		return db.Data{}, err
	}

	return d, nil
}

// Update updates records in DB, handling translations
func (n NotesService) Update(id string, data db.Data) error {
	err := n.Delete(id)
	if err != nil {
		return err
	}
	return n.Create(id, data)
}

// Delete deletes the data corresponding to id in store
// handing translations and deleting nothing
func (n NotesService) Delete(id string) error {
	hash := sha3.Sum256([]byte(id))
	q := db.Query{ID: base64.StdEncoding.EncodeToString(hash[:])}
	doc, err := q.ToBSON()
	if err != nil {
		n.logger.Sugar().Errorf("%s : %v", "[Service] [Delete] [ToBSON]", err)
		return err
	}

	res, err := n.store.Delete(doc)
	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		err = fmt.Errorf("expected deleted count to be 1, was: %d", res.DeletedCount)
		n.logger.Sugar().Errorf("%s : %v", "[Service] [Delete] [DeletedCount]", err)
		return err
	}

	return nil
}

// Exists gets data with id from DB, checks it against zero values
// and returns true if the record is non-zero
func (n NotesService) Exists(id string) (bool, error) {
	d, err := n.Get(id)
	if err != nil {
		return false, err
	}

	// If empty, data does NOT exist so NOT it
	return !d.IsEmpty(), nil
}

// NewNotesService creates a new instance of NotesService
func NewNotesService(st store.Store, lg *zap.Logger) NotesService {
	return NotesService{
		logger: lg,
		store:  st,
	}
}
