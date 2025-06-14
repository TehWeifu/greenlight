package data

import (
	"database/sql"
	"errors"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Models struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	Movies      MovieModel
	Permissions PermissionModel // Add a new Permissions field.
	Tokens      TokenModel      // Add a new Tokens field
	Users       UserModel       // Add a new Users field.
}

// For ease of use, we also add a New () method which return a Model struct containing
// the initialized MovieModel

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Permissions: PermissionModel{DB: db}, // Initialize a new PermissionsModel instance.
		Tokens:      TokenModel{DB: db},      // Initialize a new TokenModel instance.
		Users:       UserModel{DB: db},       // Initialize a new UserModel instance.
	}
}
