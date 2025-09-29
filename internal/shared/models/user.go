// Package models - collection of models
package models

import (
	"time"
)

type User struct {
	ID           int       `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"passwordHash"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	Role         UserRole  `json:"role" db:"role"`
	Active       bool      `json:"active" db:"active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type UserRole string

const (
	RoleStudent    UserRole = "student"
	RoleInstructor UserRole = "teacher"
	RoleAdmin      UserRole = "admin"
)

// IsValid - Check if role is valid
func (r UserRole) IsValid() bool {
	return r == RoleStudent || r == RoleInstructor || r == RoleAdmin
}

// String - return string representaiton of the role
func (r UserRole) String() string {
	return string(r)
}

// GetDisplayName - returns the user's display name
func (u *User) GetDisplayName() string {
	if u.FirstName != "" && u.LastName != "" {
		return u.FirstName + " " + u.LastName
	}
	return u.Username
}

// IsActive - Check if user is active
func (u *User) IsActive() bool {
	return u.Active
}

// CanAccess - CHECK if user has required role
func (u *User) CanAccess(requiredRole UserRole) bool {
	if !u.IsActive() {
		return false
	}

	if u.Role == RoleAdmin {
		return true
	}

	if u.Role == RoleInstructor && requiredRole == RoleStudent {
		return true
	}

	return u.Role == requiredRole
}

// CreateUserRequest - represent the request to create a user
type CreateUserRequest struct {
	Username  string   `json:"username" validate:"required,min=3,max=50"`
	Email     string   `json:"email" validate:"required,email"`
	Password  string   `json:"password" validate:"required,min=6"`
	FirstName string   `json:"first_name" validate:"required,min=1,max=50"`
	LastName  string   `json:"last_name" validate:"required,min=1,max=50"`
	Role      UserRole `json:"role" validate:"required"`
}

// UpdateUserRequset - represent the request to update a user
type UpdateUserRequset struct {
	Username  *string   `json:"username,omitempty" validate:"omitempty,min=3,max=50"`
	Email     *string   `json:"email,omitempty" validate:"omitempty,email"`
	Password  *string   `json:"password,omitempty" validate:"omitempty,min=6"`
	FirstName *string   `json:"first_name,omitempty" validate:"omitempty,min=1,max=50"`
	LastName  *string   `json:"last_name,omitempty" validate:"omitempty,min=1,max=50"`
	Role      *UserRole `json:"role,omitempty" validate:"omitempty"`
	Active    *bool     `json:"active,omitempty"`
}

// UserResponse - represent the response when returning user data
type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      UserRole  `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToResponse - Convert User to UserResponse
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
