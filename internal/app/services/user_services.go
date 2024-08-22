package services

import (
	"cli-project/internal/domain/interfaces"
	"cli-project/internal/domain/models"
	"cli-project/pkg/utils"
	pwd "cli-project/pkg/utils/password"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userRepo interfaces.UserRepository
	//userWG   *sync.WaitGroup
}

func NewUserService(userRepo interfaces.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		//userWG:   &sync.WaitGroup{},
	}
}

// SignUp creates a new user account
func (s *UserService) SignUp(user models.StandardUser) error {
	// Generate a new UUID for the user
	userID := utils.GenerateUUID()
	user.StandardUser.ID = userID

	// Hash the user password
	hashedPassword, err := pwd.HashPassword(user.StandardUser.Password)
	if err != nil {
		return fmt.Errorf("could not hash password: %v", err)
	}
	user.StandardUser.Password = hashedPassword

	// Set default role
	user.StandardUser.Role = "user"

	// Register the user
	err = s.userRepo.RegisterUser(user)
	if err != nil {
		return fmt.Errorf("could not register user: %v", err)
	}

	return nil
}

// Login authenticates a user
func (s *UserService) Login(username, password string) error {
	// Declare errors
	var (
		ErrInvalidCredentials = errors.New("username or password incorrect")
		ErrUserNotFound       = errors.New("user not found")
	)
	// Retrieve the user by username
	user, err := s.userRepo.FetchUser(username)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrUserNotFound // Return error if user doesn't exist
		}
		return fmt.Errorf("could not retrieve user: %v", err)
	}

	// Verify the password
	if !pwd.VerifyPassword(password, user.StandardUser.Password) {
		return ErrInvalidCredentials
	}

	return nil
}

// ViewDashboard retrieves the dashboard for the active user
func (s *UserService) ViewDashboard() error {
	// Placeholder implementation
	return nil
}

// UpdateProgress updates the user's progress in some context
func (s *UserService) UpdateProgress(userId string, progressData interface{}) error {
	// Placeholder implementation
	return nil
}

func (s *UserService) CountActiveUserInLast24Hours() (int64, error) {
	count, err := s.userRepo.CountActiveUsersInLast24Hours()
	if err != nil {
		return count, err
	}
	return count, nil
}

func (s *UserService) Logout() {
	// update lastseen of user
	// logout the user
}

func (s *UserService) IsEmailUnique(email string) (bool, error) {
	return s.userRepo.IsEmailUnique(email)
}

func (s *UserService) IsUsernameUnique(username string) (bool, error) {
	return s.userRepo.IsUsernameUnique(username)
}

func (s *UserService) IsLeetcodeIDUnique(leetcodeID string) (bool, error) {
	return s.userRepo.IsLeetcodeIDUnique(leetcodeID)
}

//func (s *UserService) WaitForCompletion() {
//	s.userWG.Wait()
//}
