package updatenote

// Request is the actual data structure for the update note request body.
// swagger:model updatenoteRequest
type Request struct {
	// ID of the note to update.
	// required: true
	ID      string `json:"id" bson:"id"`
	// New content for the note.
	// required: true
	Note    string `json:"note" bson:"note"`
	// Current password for the note.
	// required: true
	Pass    string `json:"pass" bson:"pass"`
	// Optional: New password for the note. If not provided, the password will not be changed.
	NewPass string `json:"new_pass,omitempty" bson:"new_pass,omitempty"`
}

// updatenoteParams defines the parameters for the update note endpoint.
// swagger:parameters updatenoteEndpoint
type updatenoteParams struct {
	// The note update details.
	//
	// in: body
	// required: true
	Body Request `json:"body"`
}

// Update note response body
// swagger:response updatenoteResponse
type Response struct {
	// ID of the updated note.
	ID string `json:"id" bson:"id"`
}
