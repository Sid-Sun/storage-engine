package updatepass

// Request is the actual data structure for the update password request body.
// swagger:model updatepassRequest
type Request struct {
	// ID of the note to update.
	// required: true
	ID      string `json:"id" bson:"id"`
	// Current password for the note.
	// required: true
	Pass    string `json:"pass" bson:"pass"`
	// New password for the note.
	// required: true
	NewPass string `json:"new_pass" bson:"new_pass"`
}

// updatepassParams defines the parameters for the update password endpoint.
// swagger:parameters updatepassEndpoint
type updatepassParams struct {
	// The password update details.
	//
	// in: body
	// required: true
	Body Request `json:"body"`
}

// Update password response body
// swagger:response updatepassResponse
type Response struct {
	// ID of the updated note.
	ID string `json:"id" bson:"id"`
}
