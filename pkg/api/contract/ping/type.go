package ping

// Response defines response structure for ping
type Response struct {
	Message string `json:"message" bson:"message"`
}
