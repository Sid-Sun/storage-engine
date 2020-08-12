package create

type Request struct {
	ID   string `json:"id" bson:"id"`
	Pass string `json:"pass" bson:"pass"`
	Note string `json:"note" bson:"note"`
}

type Response struct {
	ID string `json:"id" bson:"id"`
}
