// Package classification notes-api Notes API
//
// This is a sample API for managing notes.
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//	Contact: Your Name <you@example.com> https://example.com
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
	"github.com/sid-sun/notes-api/cmd/config"
	"github.com/sid-sun/notes-api/pkg/api"
	"github.com/sid-sun/notes-api/pkg/api/contract/create"
	"github.com/sid-sun/notes-api/pkg/api/contract/delete"
	"github.com/sid-sun/notes-api/pkg/api/contract/ping"
	"github.com/sid-sun/notes-api/pkg/api/contract/read"
	"github.com/sid-sun/notes-api/pkg/api/contract/updatenote"
	"github.com/sid-sun/notes-api/pkg/api/contract/updatepass"
)

// Not using these directly, but need them for swagger generation
var _ = create.Request{}
var _ = create.Response{}
var _ = read.Request{}
var _ = read.Response{}
var _ = updatenote.Request{}
var _ = updatenote.Response{}
var _ = updatepass.Request{}
var _ = updatepass.Response{}
var _ = delete.Request{}
var _ = delete.Response{}
var _ = ping.Response{}

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
