package delete

type Request struct {
	ID   string `json:"id" bson:"id"`
	Pass string `json:"pass" bson:"pass"`
}

type Response struct {
	ID string `json:"id" bson:"id"`
}
