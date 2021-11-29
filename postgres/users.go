package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/phifertoo/go_graphql_server/models"
)

type UsersRepo struct {
	DB *pg.DB
}

// GetUserById method
func (u *UsersRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
