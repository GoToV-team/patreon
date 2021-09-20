package store

import "patreon/internal/models"

type UserRepository interface {
	Create(*models.User) error
	FindByLogin(string) (*models.User, error)
}