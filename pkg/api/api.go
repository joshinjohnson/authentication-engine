package api

import (
	"github.com/joshinjohnson/authentication-engine/internal/access"
	"github.com/joshinjohnson/authentication-engine/internal/emulation"
	"github.com/joshinjohnson/authentication-engine/internal/operation"
	customErrors "github.com/joshinjohnson/authentication-engine/pkg/errors"
	"github.com/joshinjohnson/authentication-engine/pkg/models"
)

type AuthenticationEngine interface {
	// Authenticate takes in user credentials and checks it against values present in datastore
	// If found, it returns user details, else user not found error
	Authenticate(userCredential models.UserCredential) (models.UserDetails, error)

	// Register takes in user credentials and user details and saves it in datastore and returns error if there's any
	Register(userCredential models.UserCredential, userDetails models.UserDetails) error
}

// New takes in config and returns AuthenticationEngine instance
func New(cfg models.Config) (AuthenticationEngine, error) {
	switch cfg.Mode {
	case models.Emulation:
		return emulation.NewEmulationEngine()
	case models.Operation:
		return operation.NewOperationEngine(access.NewAuthenticationAccess())
	default:
		return nil, customErrors.ErrInvalidMode
	}
}
