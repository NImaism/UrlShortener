package handler

import (
	"Shorterism/internal/middleware"
	"Shorterism/internal/model"
	"Shorterism/internal/store"
	"Shorterism/internal/utility"
	"fmt"
	goaway "github.com/TwiN/go-away"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
)

type App struct {
	Store  *store.Mognodb
	Logger *zap.Logger
}

func (s *App) Register(e *gin.Context) {
	Data := new(model.User)
	if err := e.Bind(Data); err != nil {
		s.Logger.Info("handler.app.bind failed", zap.Error(err))
		e.AbortWithStatus(http.StatusBadRequest)
		return
	}
	Validator := validator.New()
	if err := Validator.Struct(Data); err != nil {
		s.Logger.Info("validator data has missing", zap.Error(err))
		e.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := s.Store.RegisterUser(*Data); err != nil {
		s.Logger.Error("store.registeruser failed", zap.Error(err))
		e.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	e.JSON(200, model.SuccessResponse(nil))
}

func (s *App) Login(e *gin.Context) {
	Data := new(model.LoginModel)
	if err := e.Bind(Data); err != nil {
		s.Logger.Info("handler.app.bind failed", zap.Error(err))
		e.AbortWithStatus(http.StatusBadRequest)
		return
	}

	Valid := validator.New()
	if err := Valid.Struct(Data); err != nil {
		s.Logger.Info("validator data has missing", zap.Error(err))
		e.AbortWithStatus(http.StatusBadRequest)
		return
	}

	Login, err := s.Store.Login(*Data)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			s.Logger.Info("store.login failed", zap.Error(err))
			e.AbortWithStatus(http.StatusNotFound)
			return
		}
		s.Logger.Error("store.login failed", zap.Error(err))
		e.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !Login {
		e.JSON(http.StatusOK, model.ErrorResponse(nil, "Account Not Found"))
		return
	}

	Token, err := utility.GenerateToken(Data.Email)
	if err != nil {
		s.Logger.Error("utility.generatetoken failed", zap.Error(err))
		e.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	e.JSON(200, model.SuccessResponse(Token))
}

func (s *App) RegisterLink(e *gin.Context) {
	TokenData, err := utility.ExtractTokenData(e)
	if err != nil {
		s.Logger.Error("utility.extractTokendata failed", zap.Error(err))
		e.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	Data := new(model.Url)
	if err := e.Bind(Data); err != nil {
		s.Logger.Info("handler.app.bind failed", zap.Error(err))
		e.AbortWithStatus(http.StatusBadRequest)
		return
	}

	Data.Author = TokenData["Email"]
	Data.Expire = time.Now().Add(24 * time.Hour).Unix()

	Valid := validator.New()
	if err := Valid.Struct(Data); err != nil {
		s.Logger.Info("validator data has missing", zap.Error(err))
		e.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if goaway.IsProfane(Data.Name) {
		s.Logger.Info("goaway.isprofane find bad word", zap.Error(err))
		e.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := s.Store.NewLink(*Data); err != nil {
		if err.Error() == "one Link is active with this Name" {
			s.Logger.Info("store.newlink has duplicate")
			e.AbortWithStatus(http.StatusBadRequest)
			return
		}
		s.Logger.Error("store.newlink failed", zap.Error(err))
		e.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	s.Logger.Info(fmt.Sprintf("\n [NewLink] Author : %s | Name : %s | Url : %s \n", TokenData["Email"], Data.Name, Data.Url))

	e.JSON(200, model.SuccessResponse("localhost:8080/"+Data.Name))
}

func (s *App) Link(e *gin.Context) {
	Data, Err := s.Store.Link(e.Param("link"))
	if Err != nil {
		if Err.Error() == "mongo: no documents in result" {
			s.Logger.Info("store.find Cant Find Link", zap.String("Link", e.Param("link")))
			e.AbortWithStatus(http.StatusNotFound)
			return
		}
		s.Logger.Error("store.find Cant Find Link", zap.Error(Err))
		e.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	e.Redirect(http.StatusTemporaryRedirect, Data.Url)
}

func (s *App) SetRouting(e *gin.Engine) {
	e.GET("/:link", s.Link)

	Api := e.Group("/api/v1/").Use(middleware.JwtAuthMiddleware())
	Api.POST("CreateLink/", s.RegisterLink)

	Account := e.Group("/account/v1/")
	Account.POST("Register/", s.Register)
	Account.POST("Login/", s.Login)
}
