package model

type Item struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Price int32  `json:"price,omitempty" bson:"price,omitempty"`
}
