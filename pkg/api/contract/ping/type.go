package ping

// Ping response body
// swagger:response pingResponse
type Response struct {
	// The ping message
	Message string `json:"message" bson:"message"`
}
