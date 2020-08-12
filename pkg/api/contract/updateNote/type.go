package updateNote

type Request struct {
	ID      string `json:"id" bson:"id"`
	Note    string `json:"note" bson:"note"`
	Pass    string `json:"pass" bson:"pass"`
	NewPass string `json:"new_pass" bson:"new_pass"`
}

type Response struct {
	ID string `json:"id" bson:"id"`
}
