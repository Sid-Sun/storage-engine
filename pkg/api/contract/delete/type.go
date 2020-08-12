package delete

// Request defines request structure for delete
type Request struct {
	ID   string `json:"id" bson:"id"`
	Pass string `json:"pass" bson:"pass"`
}

// Response defines response structure for delete
type Response struct {
	ID string `json:"id" bson:"id"`
}
