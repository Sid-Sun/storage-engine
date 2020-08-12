package db

type Data struct {
	AAD  []byte   `json:"aad" bson:"aad"`
	Hash [32]byte `json:"hash" bson:"hash"`
	Note []byte   `json:"note" bson:"note"`
}

func NewDataInstance(aad []byte, hash [32]byte, note []byte) Data {
	return Data{
		AAD:  aad,
		Hash: hash,
		Note: note,
	}
}
