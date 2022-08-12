package Service

import (
	"Shorterism/Model"
	"Shorterism/Repository"
)

type LinkService interface {
	CreateLink(Data Model.Url) error
	Link(Name string) (Model.Url, error)
}

type linkService struct{}

func NewShorterService() LinkService {
	return linkService{}
}

func (linkService) CreateLink(Data Model.Url) error {
	Srv := Repository.NewLinkShorter()
	return Srv.NewLink(Data)
}

func (linkService) Link(Name string) (Model.Url, error) {
	Srv := Repository.NewLinkShorter()
	return Srv.Link(Name)
}
