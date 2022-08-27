package Model

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	Email string
	jwt.StandardClaims
}

type LoginModel struct {
	Email string `json:"Email" validate:"required"`
	Pass  string `json:"Pass" validate:"required"`
}

type User struct {
	Id       string `bson:"_id,omitempty"`
	UserName string `bson:"UserName" json:"UserName" validate:"required"`
	Email    string `bson:"Email" json:"Email" validate:"required"`
	Pass     string `bson:"Pass" json:"Pass" validate:"required"`
}
