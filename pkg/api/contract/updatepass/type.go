package updatepass

// Request defines request structure for updatePass
type Request struct {
	ID      string `json:"id" bson:"id"`
	Pass    string `json:"pass" bson:"pass"`
	NewPass string `json:"new_pass" bson:"pass"`
}

// Response defines response structure for updatePass
type Response struct {
	ID string `json:"id" bson:"id"`
}
