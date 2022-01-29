package operation

import (
	"github.com/joshinjohnson/authentication-service/internal/access"
	"github.com/joshinjohnson/authentication-service/pkg/models"
	"reflect"
	"testing"
)

func TestEngine_Authenticate(t *testing.T) {
	accessWithData := access.NewAuthenticationAccess()
	accessWithData.StoreUser(models.UserCredential{
		Username: "username",
		Email:    "email@website.com",
		Password: "pass",
	}, models.UserDetails{
		FirstName:   "firstname",
		LastName:    "lastname",
	})
	type args struct {
		userCredential models.UserCredential
	}
	tests := []struct {
		name    string
		access *access.AuthenticationAccess
		args    args
		want    models.UserDetails
		wantErr bool
	}{
		{
			name:    "negative",
			access: nil,
			args:    args{},
			want:    models.UserDetails{},
			wantErr: true,
		},
		{
			name:    "positive",
			access:  accessWithData,
			args:    args{},
			want:    models.UserDetails{
				FirstName:   "firstname",
				LastName:    "lastname",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Engine{
				access: tt.access,
			}
			got, err := e.Authenticate(tt.args.userCredential)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEngine_Register(t *testing.T) {
	userDetails := models.UserDetails{
		FirstName:   "firstname",
		LastName:    "lastname",
	}
	type fields struct {
		access *access.AuthenticationAccess
	}
	type args struct {
		userCredential models.UserCredential
		userDetails    models.UserDetails
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		userDetails models.UserDetails
	}{
		{
			name:    "negative",
			fields:  fields{},
			args:    args{},
			wantErr: true,
		},
		{
			name:        "positive",
			fields:      fields{
				access: access.NewAuthenticationAccess(),
			},
			args:        args{
				userCredential: models.UserCredential{
					Username: "username",
					Email:    "",
					Password: "123asd",
				},
				userDetails:    userDetails,
			},
			wantErr:     false,
			userDetails: userDetails,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Engine{
				access: tt.fields.access,
			}
			if err := e.Register(tt.args.userCredential, tt.args.userDetails); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got, err := e.access.FetchUserDetails(models.UserCredential{}); err != nil || !reflect.DeepEqual(got, tt.userDetails) {
				t.Errorf("Register() want = %v, got %v", tt.userDetails, got)
			}
		})
	}
}

func TestNewOperationEngine(t *testing.T) {
	type args struct {
		access *access.AuthenticationAccess
	}
	tests := []struct {
		name    string
		args    args
		want    *Engine
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOperationEngine(tt.args.access)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOperationEngine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOperationEngine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMD5Hash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive",
			args: args{password: "pass"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMD5Hash(tt.args.password); got == tt.args.password {
				t.Errorf("getMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}