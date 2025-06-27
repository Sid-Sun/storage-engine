package delete

// Request is the actual data structure for the delete request body.
// swagger:model deleteRequest
type Request struct {
	// ID of the note to delete.
	// required: true
	ID   string `json:"id" bson:"id"`
	// Password for the note.
	// required: true
	Pass string `json:"pass" bson:"pass"`
}

// deleteNoteParams defines the parameters for the delete note endpoint.
// swagger:parameters deleteNoteEndpoint
type deleteNoteParams struct {
	// The note ID and password to delete.
	//
	// in: body
	// required: true
	Body Request `json:"body"`
}

// Delete response body
// swagger:response deleteResponse
type Response struct {
	// ID of the deleted note.
	ID string `json:"id" bson:"id"`
}
