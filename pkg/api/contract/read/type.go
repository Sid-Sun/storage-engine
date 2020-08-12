package read

type Request struct {
	ID   string `json:"id" bson:"id"`
	Pass string `json:"pass" bson:"id"`
}

type Response struct {
	ID   string `json:"id" bson:"id"`
	Note string `json:"note" bson:"note"`
}
