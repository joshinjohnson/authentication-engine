package emulation

import (
	"github.com/joshinjohnson/authentication-engine/pkg/models"
)

type Engine struct {
}

func NewEmulationEngine() (*Engine, error) {
	return &Engine{}, nil
}

func (e Engine) Authenticate(_ models.UserCredential) (models.UserDetails, error) {
	return models.UserDetails{}, nil
}

func (e Engine) Register(_ models.UserCredential, _ models.UserDetails) error {
	return nil
}
