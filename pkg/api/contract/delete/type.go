package delete

// DeleteRequest is the actual data structure for the delete request body.
// swagger:model deleteRequest
type DeleteRequest struct {
	// ID of the note to delete.
	// required: true
	ID string `json:"id" bson:"id"`
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
	Body DeleteRequest `json:"body"`
}

// deleteNoteResponse defines the response of delete endpoint
// swagger:response deleteResponse
type deleteNoteResponse struct {
	// in: body
	Body DeleteResponse `json:"body"`
}

// DeleteResponse is the actual response body for delete
type DeleteResponse struct {
	// ID of the deleted note.
	ID string `json:"id" bson:"id"`
}
