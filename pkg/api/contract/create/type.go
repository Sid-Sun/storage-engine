package create

// Request defines request structure for create
type Request struct {
	ID   string `json:"id" bson:"id"`
	Pass string `json:"pass" bson:"pass"`
	Note string `json:"note" bson:"note"`
}

// Response defines response structure for create
type Response struct {
	ID string `json:"id" bson:"id"`
}
