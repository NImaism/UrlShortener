package Controllers

import (
	"Shorterism/Model"
	"Shorterism/Service"
	"Shorterism/Utility"
	"fmt"
	goaway "github.com/TwiN/go-away"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
)

func Register(e *gin.Context) {
	Data := new(Model.User)
	if err := e.Bind(Data); err != nil {
		e.JSON(500, Model.ErrorResponse(err.Error(), "Error In Binding Data"))
		return
	}
	Validator := validator.New()
	if err := Validator.Struct(Data); err != nil {
		e.JSON(400, Model.ErrorResponse(err.Error(), "Data Not Valid"))
		return
	}

	Srv := Service.NewUserService()
	if err := Srv.RegisterUser(*Data); err != nil {
		e.JSON(500, Model.ErrorResponse(err.Error(), "Error In Register User"))
		return
	}

	e.JSON(200, Model.SuccessResponse(nil))
}

func Login(e *gin.Context) {
	Data := new(Model.LoginModel)
	if err := e.Bind(Data); err != nil {
		e.JSON(500, Model.ErrorResponse(err.Error(), "Error In Binding Data"))
		return
	}

	Valid := validator.New()
	if err := Valid.Struct(Data); err != nil {
		e.JSON(400, Model.ErrorResponse(err.Error(), "Data Not Valid"))
		return
	}

	Srv := Service.NewUserService()
	Login, err := Srv.Login(*Data)
	if err != nil {
		e.JSON(500, Model.ErrorResponse(err.Error(), "Error In Login User"))
		return
	}

	if !Login {
		e.JSON(http.StatusOK, Model.ErrorResponse(nil, "Account Not Found"))
		return
	}

	Token, err := Utility.GenerateToken(Data.Email)
	if err != nil {
		e.JSON(500, Model.ErrorResponse(err.Error(), "Error In Create Token"))
		return
	}

	e.JSON(200, Model.SuccessResponse(Token))
}

func RegisterLink(e *gin.Context) {
	TokenData, err := Utility.ExtractTokenData(e)
	if err != nil {
		e.JSON(500, Model.ErrorResponse(err.Error(), "Error In Extract Token Data"))
		return
	}

	Data := new(Model.Url)
	if err := e.Bind(Data); err != nil {
		e.JSON(500, Model.ErrorResponse(err.Error(), "Error In Binding Data"))
		return
	}

	Data.Author = TokenData["Email"]
	Data.Expire = time.Now().Add(24 * time.Hour).Unix()

	Valid := validator.New()
	if err := Valid.Struct(Data); err != nil {
		e.JSON(400, Model.ErrorResponse(err.Error(), "Data Not Valid"))
		return
	}

	if goaway.IsProfane(Data.Name) {
		e.JSON(403, Model.ErrorResponse(nil, "You Cant Use Bad Words"))
		return
	}

	Srv := Service.NewShorterService()
	if err := Srv.CreateLink(*Data); err != nil {
		e.JSON(400, Model.ErrorResponse(err.Error(), "Error In Create Link"))
		return
	}
	fmt.Printf("\n [NewLink] Author : %s | Name : %s | Url : %s \n", TokenData["Email"], Data.Name, Data.Url)
	e.JSON(200, "localhost:8080/"+Data.Name)
}

func Link(e *gin.Context) {
	Srv := Service.NewShorterService()
	Data, Err := Srv.Link(e.Param("link"))
	if Err != nil {
		e.JSON(400, Model.ErrorResponse(Err.Error(), "Error In Find Link"))
		return
	}
	e.Redirect(http.StatusTemporaryRedirect, Data.Url)
}
