package models

type Page struct {
	ID  int    `json:"page" bson:"Page_No"`
	Key string `json:"key" bson:"Contents"`
}
