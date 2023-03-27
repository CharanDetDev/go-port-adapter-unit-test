package repository

import "github.com/CharanDetDev/go-port-adapter-unit-test/domain"

type personRepo struct{}

func NewPersonRepo() domain.PersonRepo {
	return &personRepo{}
}
