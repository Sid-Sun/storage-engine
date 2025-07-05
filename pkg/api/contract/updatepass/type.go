package updatepass

// UpdatePassRequest is the actual data structure for the update password request body.
// swagger:model updatepassRequest
type UpdatePassRequest struct {
	// ID of the note to update.
	// required: true
	ID string `json:"id" bson:"id"`
	// Current password for the note.
	// required: true
	Pass string `json:"pass" bson:"pass"`
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
	Body UpdatePassRequest `json:"body"`
}

// updatepassNoteResponse defines the response of updatepass endpoint
// swagger:response updatepassResponse
type updatepassNoteResponse struct {
	// in: body
	Body UpdatePassResponse `json:"body"`
}

// UpdatePassResponse is the actual response body for updatepass
type UpdatePassResponse struct {
	// ID of the updated note.
	ID string `json:"id" bson:"id"`
}
