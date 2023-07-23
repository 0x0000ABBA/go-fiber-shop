package services

import (
	"errors"
	"fiber-shop/internal/data/models"

	"golang.org/x/exp/slices"
	// "fiber-shop/internal/data/models"
	// "fiber-shop/pkg/pgsql"
	// "github.com/jackc/pgx"
)

var admin1 = models.User{Id: 1, FirstName: "admin", LastName: "admin", Email: "admin@admin.go", Password: "adminpword"}
var admin2 = models.User{Id: 2, FirstName: "admin2", LastName: "admin2", Email: "admin2@admin.go", Password: "adminpword2"}

var user1 = models.User{Id: 3, FirstName: "user", LastName: "user", Email: "user@admin.go", Password: "userpword"}
var user2 = models.User{Id: 4, FirstName: "user2", LastName: "user2", Email: "user2@admin.go", Password: "userpword2"}

var users []models.User = []models.User{admin1, admin2, user1, user2}

// TODO implement me
func GetAllUsers() ([]models.User, error) {
	return users, nil
}

func GetUserById(id int) (models.User, error) {

	index := slices.IndexFunc[models.User](users, func(u models.User) bool { return u.Id == id})

	if index == -1 {
		return models.User{}, errors.New("user not found")
	}

	user := users[index]

	return user, nil
}
