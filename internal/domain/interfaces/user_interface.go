package interfaces

import (
	"cli-project/internal/domain/models"
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(*models.StandardUser) error
	UpdateUserProgress(userID uuid.UUID, questionID []string) error
	FetchAllUsers() (*[]models.StandardUser, error)
	FetchUserByID(string) (*models.StandardUser, error)
	FetchUserByUsername(context.Context, string) (*models.StandardUser, error)
	FetchUserProgress(string) (*[]string, error)
	UpdateUserDetails(*models.StandardUser) error
	BanUser(string) error
	UnbanUser(string) error
	CountActiveUsersInLast24Hours() (int, error)
	IsUsernameUnique(string) (bool, error)
	IsEmailUnique(string) (bool, error)
	IsLeetcodeIDUnique(string) (bool, error)
}
