package repositories

import (
	"go-api/pkg/utils"
)

func FetchAllUsers() ([]utils.Users, error) {
	return []utils.Users{
		{ID: 1, Name: "John sign"},
		{ID: 2, Name: "John path"},
	}, nil
}
