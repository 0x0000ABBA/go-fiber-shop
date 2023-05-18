package services

import (
	"errors"
	"fiber-shop/internal/data/models"
	// "fiber-shop/internal/data/models"
	// "fiber-shop/pkg/pgsql"
	// "github.com/jackc/pgx"
)

// TODO implement me
func GetAllUsers() ([]models.User, error) {

	user1 := models.User{Id: 1, FirstName: "admin", LastName: "admin", Email: "admin@admin.go", Password: "adminpword"}
	user2 := models.User{Id: 2, FirstName: "admin2", LastName: "admin2", Email: "admin2@admin.go", Password: "adminpword2"}

	data := []models.User{user1, user2}

	return data, nil
}

func GetUserById(id string) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}
