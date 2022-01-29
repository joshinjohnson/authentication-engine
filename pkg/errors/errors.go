package errors

import "fmt"

var (
	ErrEmptyAccess = fmt.Errorf("empty access received")
	ErrUserNotFound = fmt.Errorf("no user found")
	ErrInvalidMode = fmt.Errorf("invalid mode received")
)