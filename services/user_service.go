package services

import (
	"go-api/pkg/utils"
	"go-api/repositories"
)

func GetUsers() ([]utils.Users, error) {
	return repositories.FetchAllUsers()
}
