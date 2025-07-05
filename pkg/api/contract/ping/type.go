package ping

// Ping response body
// swagger:response pingResponse
type PingResponse struct {
	// The ping message
	Message string `json:"message" bson:"message"`
}
