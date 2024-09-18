package repo

import (
	"database/sql"
	"fintrack/internal/domain/entity"
	"time"
)

// UserRepository defines the interface for user repository methods
type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByID(id uint64) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint64) error
	GetUserByEmail(email string) (*entity.User, error)       // Method for fetching user by email
	GetUserByUserName(userName string) (*entity.User, error) // New method for fetching user by username
}

// UserRepo implements UserRepository for working with user data
type UserRepo struct {
	DB *sql.DB
}

// NewUserRepo creates a new UserRepo
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// CreateUser inserts a new user into the database
func (repo *UserRepo) CreateUser(user *entity.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	query := `INSERT INTO users (user_name, first_name, last_name, email, password, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := repo.DB.Exec(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByEmail fetches a user by their email
func (repo *UserRepo) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, user_name, first_name, last_name, email, password, status, created_at, updated_at FROM users WHERE email = ?`
	err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return user, nil
}

// GetUserByUserName fetches a user by their username
func (repo *UserRepo) GetUserByUserName(userName string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, user_name, first_name, last_name, email, password, status, created_at, updated_at FROM users WHERE user_name = ?`
	err := repo.DB.QueryRow(query, userName).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return user, nil
}

// GetUserByID fetches a user by their ID
func (repo *UserRepo) GetUserByID(id uint64) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, user_name, first_name, last_name, email, status, created_at, updated_at FROM users WHERE id = ?`
	err := repo.DB.QueryRow(query, id).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user's details
func (repo *UserRepo) UpdateUser(user *entity.User) error {
	user.UpdatedAt = time.Now()
	query := `UPDATE users SET user_name = ?, first_name = ?, last_name = ?, email = ?, password = ?, status = ?, updated_at = ? WHERE id = ?`
	_, err := repo.DB.Exec(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser removes a user from the database
func (repo *UserRepo) DeleteUser(id uint64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := repo.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
