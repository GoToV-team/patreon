package repository_like

import (
	"patreon/internal/app/models"
)

const NoAwards = -1

type Repository interface {
	// Create Errors:
	// 		app.GeneralError with Errors
	// 			repository.DefaultErrDB
	Create(post *models.Post) (int64, error)

	// GetPost Errors:
	//		repository.NotFound
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	GetPost(postID int64) (*models.Post, error)

	// GetPosts Errors:
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	GetPosts(page int64) ([]models.Creator, error)

	// UpdatePost Errors:
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	UpdatePost(post *models.Post) error

	// UpdateCoverPost Errors:
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	UpdateCoverPost(postId int64, cover string) error

	// Delete Errors:
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	Delete(postId int64) error
}
