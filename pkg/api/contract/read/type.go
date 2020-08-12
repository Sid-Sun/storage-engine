package read

// Request defines request structure for read
type Request struct {
	ID   string `json:"id" bson:"id"`
	Pass string `json:"pass" bson:"id"`
}

// Response defines response structure for read
type Response struct {
	ID   string `json:"id" bson:"id"`
	Note string `json:"note" bson:"note"`
}
