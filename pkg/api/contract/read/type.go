package read

// ReadRequest is the actual data structure for the read request body.
// swagger:model readRequest
type ReadRequest struct {
	// ID of the note to read.
	// required: true
	ID string `json:"id" bson:"id"`
	// Password for the note.
	// required: true
	Pass string `json:"pass" bson:"pass"` // Corrected bson tag here
}

// readNoteParams defines the parameters for the read note endpoint.
// swagger:parameters readNoteEndpoint
type readNoteParams struct {
	// The note ID and password to read.
	//
	// in: body
	// required: true
	Body ReadRequest `json:"body"`
}

// readNoteResponse defines the response of read endpoint
// swagger:response readResponse
type readNoteResponse struct {
	// in: body
	Body ReadResponse `json:"body"`
}

// ReadResponse is the actual response body for read
type ReadResponse struct {
	// ID of the note.
	ID string `json:"id" bson:"id"`
	// Content of the note.
	Note string `json:"note" bson:"note"`
}
