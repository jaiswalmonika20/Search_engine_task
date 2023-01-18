package models

type Page struct {
	ID  int    `json:"id" bson:"Page_No"`
	Key string `json:"key" bson:"Contents"`
}
