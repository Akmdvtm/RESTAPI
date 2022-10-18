package service

import "restapi/pkg/repository"

type Authorization interface {
}

type UsersList interface {
}

type Service struct {
	Authorization
	UsersList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
