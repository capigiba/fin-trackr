package user

import (
	"errors"
	"fintrack/internal/domain/entity"
	"fintrack/internal/pkg/reason"
	"fintrack/internal/repo/user"
	"fintrack/internal/service/auth"

	"golang.org/x/crypto/bcrypt"
)

// UserService handles user-related business logic
type UserService struct {
	userRepo    user.UserRepository
	authService *auth.AuthService
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo user.UserRepository, authService *auth.AuthService) *UserService {
	return &UserService{
		userRepo:    userRepo,
		authService: authService,
	}
}

// RegisterUser handles the registration of a new user
func (s *UserService) RegisterUser(userName, firstName, lastName, email, password string) error {
	// Check if the email is already registered
	existingUser, _ := s.userRepo.GetUserByEmail(email)
	if existingUser != nil {
		return errors.New(reason.EmailAlreadyRegistered.Message())
	}

	// Check if the username is already registered
	existingUserByUserName, _ := s.userRepo.GetUserByUserName(userName)
	if existingUserByUserName != nil {
		return errors.New(reason.UserNameAlreadyRegistered.Message())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword),
	}

	return s.userRepo.CreateUser(user)
}

// Login authenticates the user and returns a JWT token
func (s *UserService) Login(userName, password string) (string, error) {
	// Check if user exists by username
	user, err := s.userRepo.GetUserByUserName(userName)
	if err != nil || user == nil {
		return "", errors.New(reason.UserNotFound.Message())
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New(reason.InvalidCredentials.Message())
	}

	return s.authService.GenerateToken(user) // Assuming authService generates the token
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id uint64) (*entity.User, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user's details
func (s *UserService) UpdateUser(id uint64, userName, firstName, lastName, email, password string, status int) error {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	user.UserName = userName
	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	user.Status = status

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return s.userRepo.UpdateUser(user)
}

// DeleteUser removes a user
func (s *UserService) DeleteUser(id uint64) error {
	return s.userRepo.DeleteUser(id)
}
