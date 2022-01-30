package operation

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/joshinjohnson/authentication-engine/internal/access"
	customErrors "github.com/joshinjohnson/authentication-engine/pkg/errors"
	"github.com/joshinjohnson/authentication-engine/pkg/models"
)

type Engine struct {
	access *access.AuthenticationAccess
}

func NewOperationEngine(access *access.AuthenticationAccess) (*Engine, error) {
	if access == nil {
		return nil, customErrors.ErrEmptyAccess
	}

	return &Engine{access: access}, nil
}

func (e Engine) Authenticate(userCredential models.UserCredential) (models.UserDetails, error) {
	if e.access == nil {
		return models.UserDetails{}, customErrors.ErrEmptyAccess
	}
	return e.access.FetchUserDetails(userCredential)
}

func (e Engine) Register(userCredential models.UserCredential, userDetails models.UserDetails) error {
	if e.access == nil {
		return customErrors.ErrEmptyAccess
	}
	userCredential.Password = getMD5Hash(userCredential.Password)
	return e.access.StoreUser(userCredential, userDetails)
}

func getMD5Hash(password string) string {
	b := md5.Sum([]byte(password))
	return hex.EncodeToString(b[:])
}
