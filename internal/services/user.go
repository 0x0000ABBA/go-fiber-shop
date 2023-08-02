package services

import (
	"errors"
	"fiber-shop/internal/data/models"

	"golang.org/x/exp/slices" // TODO remove later
)

var admin1 = models.User{Id: "8450c898-b2ff-4edd-927d-d9e9c620f871", FirstName: "admin", LastName: "admin", Email: "admin@admin.go", Password: "adminpword"}
var admin2 = models.User{Id: "8be8e038-de9e-483a-9fb7-bf82a8dbc0a4", FirstName: "admin2", LastName: "admin2", Email: "admin2@admin.go", Password: "adminpword2"}

var user1 = models.User{Id: "5ddeb299-4de7-4f81-a43d-fe2be3301dbc", FirstName: "user", LastName: "user", Email: "user@admin.go", Password: "userpword"}
var user2 = models.User{Id: "27e4857f-e67a-49de-b228-f10409f0f745", FirstName: "user2", LastName: "user2", Email: "user2@admin.go", Password: "userpword2"}

var users []models.User = []models.User{admin1, admin2, user1, user2}

// TODO implement me
func GetAllUsers() ([]models.User, error) {
	return users, nil
}

func GetUserById(id string) (models.User, error) {

	index := slices.IndexFunc[models.User](users, func(u models.User) bool { return u.Id == id })

	if index == -1 {
		return models.User{}, errors.New("user not found")
	}

	user := users[index]

	return user, nil
}

func UserExists(email string) bool {
	index := slices.IndexFunc[models.User](users, func(u models.User) bool { return u.Email == email }) // FIXME

	return index != -1
}
