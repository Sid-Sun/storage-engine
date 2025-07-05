// Package classification Storage Engine
//
// An API which can be easily integrated into applications to provide state/storage.
//
//	Schemes: http
//	Host: localhost:5050
//	BasePath: /
//	Version: 1.0.0
//	Contact: Sid Sun <sid@sidsun.com> https://sidsun.com
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"github.com/sid-sun/storage-engine/cmd/config"
	"github.com/sid-sun/storage-engine/pkg/api"
	"github.com/sid-sun/storage-engine/pkg/api/contract/create"
	"github.com/sid-sun/storage-engine/pkg/api/contract/delete"
	"github.com/sid-sun/storage-engine/pkg/api/contract/ping"
	"github.com/sid-sun/storage-engine/pkg/api/contract/read"
	"github.com/sid-sun/storage-engine/pkg/api/contract/updatenote"
	"github.com/sid-sun/storage-engine/pkg/api/contract/updatepass"
)

// Not using these directly, but need them for swagger generation
var _ = create.CreateRequest{}
var _ = create.CreateResponse{}
var _ = read.ReadRequest{}
var _ = read.ReadResponse{}
var _ = updatenote.UpdateNoteRequest{}
var _ = updatenote.UpdateNoteResponse{}
var _ = updatepass.UpdatePassRequest{}
var _ = updatepass.UpdatePassResponse{}
var _ = delete.DeleteRequest{}
var _ = delete.DeleteResponse{}
var _ = ping.PingResponse{}

// ErrorMessage represents a generic error message.
// swagger:model
type ErrorMessage struct {
	Message string `json:"message"`
}

// SuccessMessage represents a generic success message.
// swagger:model
type SuccessMessage struct {
	Message string `json:"message"`
}

// A generic error response.
// swagger:response genericError
type GenericError struct {
	// in:body
	Body ErrorMessage
}

// A generic success response.
// swagger:response genericSuccess
type GenericSuccess struct {
	// in:body
	Body SuccessMessage
}

func main() {
	cfg := config.Load()
	initLogger(cfg.GetEnv())
	api.StartServer(cfg, logger)
}
