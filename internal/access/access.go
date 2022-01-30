package access

import (
    customErrors "github.com/joshinjohnson/authentication-engine/pkg/errors"
    "github.com/joshinjohnson/authentication-engine/pkg/models"
    "time"
)

const tableSize = 100

var (
    access *AuthenticationAccess
)

// AuthenticationAccess is the access layer which hides internal data source.
// local data store is used here and fields represent tables in a database
type AuthenticationAccess struct {
    credentials            []credential
    users                  []user
    credentialToUserLookup []credentialToUserLookup
}

// NewAuthenticationAccess returns an instance of AuthenticationAccess
func NewAuthenticationAccess() *AuthenticationAccess {
    if access != nil {
        return access
    }
    return &AuthenticationAccess{
        credentials:            make([]credential, 0, tableSize),
        users:                  make([]user, 0, tableSize),
        credentialToUserLookup: make([]credentialToUserLookup, 0, tableSize),
    }
}

// StoreUser takes in user credentials and user details and store it in internal datastore. If there's any error, return it
func (a *AuthenticationAccess) StoreUser(c models.UserCredential, d models.UserDetails) error {
    if a.users == nil || a.credentials == nil || a.credentialToUserLookup == nil {
        return customErrors.ErrEmptyAccess
    }
    credID, userID := int64(len(a.credentials)+1), int64(len(a.users)+1)
    a.credentials = append(a.credentials, credential{
        id:              credID,
        username:        c.Username,
        passwordHash:    c.Password,
        email:           c.Email,
        lastUpdatedTime: time.Now(),
        createdTime:     time.Now(),
    })
    a.users = append(a.users, user{
        userID:          userID,
        firstName:       d.FirstName,
        lastName:        d.LastName,
        dateOfBirth:     d.DateOfBirth,
        lastUpdatedTime: time.Now(),
        createdTime:     time.Now(),
    })
    a.credentialToUserLookup = append(a.credentialToUserLookup, credentialToUserLookup{
        credentialID: credID,
        userID:       userID,
    })
    return nil
}

// FetchUserDetails takes in userCredential models and checks if it exists in datastore and returns userDetails, else throws user not found error
func (a *AuthenticationAccess) FetchUserDetails(userCredential models.UserCredential) (models.UserDetails, error) {
    if a.users == nil || a.credentials == nil || a.credentialToUserLookup == nil {
        return models.UserDetails{}, customErrors.ErrEmptyAccess
    }
    var credID, userID int64
    for _, c := range a.credentials {
        if verifyPassword(c.passwordHash, userCredential.Password) && c.email == userCredential.Email {
            credID = c.id
            break
        }
    }
    if credID == 0 {
        return models.UserDetails{}, customErrors.ErrUserNotFound
    }
    for _, l := range a.credentialToUserLookup {
        if l.credentialID == credID {
            userID = l.userID
            break
        }
    }
    if userID == 0 {
        return models.UserDetails{}, customErrors.ErrUserNotFound
    }
    var u user
    for _, u = range a.users {
        if u.userID == userID {
            break
        }
    }
    if u.userID == 0 {
        return models.UserDetails{}, customErrors.ErrUserNotFound
    }
    return models.UserDetails{
        FirstName:   u.firstName,
        LastName:    u.lastName,
        DateOfBirth: u.dateOfBirth,
    }, nil
}

func verifyPassword(hash string, password string) bool {
    return hash == GetMD5Hash(password)
}
