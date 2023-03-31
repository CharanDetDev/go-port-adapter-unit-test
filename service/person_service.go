package service

import (
	"github.com/CharanDetDev/go-port-adapter-unit-test/domain"
)

type personService struct {
	PersonRepo     domain.PersonRepo
	RedisCacheRepo domain.RedisCacheRepo
}

func NewPersonService(personRepo domain.PersonRepo, redisCacheRepo domain.RedisCacheRepo) domain.PersonService {
	return &personService{
		PersonRepo:     personRepo,
		RedisCacheRepo: redisCacheRepo,
	}
}
