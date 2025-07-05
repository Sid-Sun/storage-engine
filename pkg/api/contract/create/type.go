package create

// CreateRequest is the actual data structure for the create request body.
// swagger:model createRequest
type CreateRequest struct {
	// Optional: ID for the note. If not provided, a random ID will be generated.
	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// Password for the note.
	// required: true
	Pass string `json:"pass" bson:"pass"`
	// Content of the note.
	// required: true
	Note string `json:"note" bson:"note"`
}

// createNoteParams defines the parameters for the create note endpoint.
// swagger:parameters createNoteEndpoint
type createNoteParams struct {
	// The note to create.
	//
	// in: body
	// required: true
	Body CreateRequest `json:"body"`
}

// createNoteResponse defines the response of create endpoint
// swagger:response createResponse
type createNoteResponse struct {
	// in: body
	Body CreateResponse `json:"body"`
}

// CreateResponse is the actual response body for create
type CreateResponse struct {
	// ID of the created note.
	ID string `json:"id" bson:"id"`
}
