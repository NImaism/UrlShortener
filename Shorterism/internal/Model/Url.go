package Model

type Url struct {
	Name   string      `bson:"Name" json:"Name" validate:"required"`
	Url    string      `bson:"Url" json:"Url" validate:"required"`
	Author interface{} `bson:"Author" validate:"required"`
	Expire int64       `bson:"Expire"`
}
