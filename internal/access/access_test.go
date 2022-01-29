package access

import (
	"github.com/joshinjohnson/authentication-service/pkg/models"
	"reflect"
	"testing"
	"time"
)

func TestAuthenticationAccess_FetchUserDetails(t *testing.T) {
	type fields struct {
		credentials            []credential
		users                  []user
		credentialToUserLookup []credentialToUserLookup
	}
	type args struct {
		userCredential models.UserCredential
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.UserDetails
		wantErr bool
	}{
		{
			name:    "negative",
			fields:  fields{},
			args:    args{},
			want:    models.UserDetails{},
			wantErr: true,
		},
		{
			name:    "negative",
			fields:  fields{
				credentials:            make([]credential, 1),
				users:                  make([]user, 1),
				credentialToUserLookup: make([]credentialToUserLookup, 1),
			},
			args:    args{userCredential: models.UserCredential{
				Username: "username",
				Email:    "",
				Password: "123asd",
			}},
			want:    models.UserDetails{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AuthenticationAccess{
				credentials:            tt.fields.credentials,
				users:                  tt.fields.users,
				credentialToUserLookup: tt.fields.credentialToUserLookup,
			}
			got, err := a.FetchUserDetails(tt.args.userCredential)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchUserDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUserDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthenticationAccess_StoreUser(t *testing.T) {
	type fields struct {
		credentials            []credential
		users                  []user
		credentialToUserLookup []credentialToUserLookup
	}
	type args struct {
		c models.UserCredential
		d models.UserDetails
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		userData models.UserDetails
	}{
		{
			name:     "negative",
			fields:   fields{},
			args:     args{},
			wantErr:  true,
			userData: models.UserDetails{},
		},
		{
			name:    "positive",
			fields:  fields{
				credentials:            make([]credential, 0, 1),
				users:                  make([]user, 0, 1),
				credentialToUserLookup: make([]credentialToUserLookup, 0, 1),
			},
			args:    args{
				c: models.UserCredential{
					Username: "username",
					Email:    "email@website.com",
					Password: "123asd",
				},
				d: models.UserDetails{
					FirstName:   "firstname",
					LastName:    "lastname",
					DateOfBirth: time.Date(1999, 05, 20, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
			userData: models.UserDetails{
				FirstName:   "firstname",
				LastName:    "lastname",
				DateOfBirth: time.Date(1999, 05, 20, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AuthenticationAccess{
				credentials:            tt.fields.credentials,
				users:                  tt.fields.users,
				credentialToUserLookup: tt.fields.credentialToUserLookup,
			}
			if err := a.StoreUser(tt.args.c, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("StoreUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if u, err := a.FetchUserDetails(models.UserCredential{
				Username: "username",
				Email:    "email@website.com",
				Password: "123asd",
			}); err != nil {
				t.Errorf("FetchUserDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if u != tt.userData {
				t.Errorf("FetchUserDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}