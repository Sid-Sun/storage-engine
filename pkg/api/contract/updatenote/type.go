package updatenote

// Request defines request structure for updateNote
type Request struct {
	ID      string `json:"id" bson:"id"`
	Note    string `json:"note" bson:"note"`
	Pass    string `json:"pass" bson:"pass"`
	NewPass string `json:"new_pass" bson:"new_pass"`
}

// Response defines response structure for updateNote
type Response struct {
	ID string `json:"id" bson:"id"`
}
