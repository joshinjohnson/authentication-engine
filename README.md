# authentication-engine
## API methods
- Authenticate(userCredential models.UserCredential) (models.UserDetails, error)
- Register(userCredential models.UserCredential, userDetails models.UserDetails) error

## Installation
- go get -u github.com/joshinjohnson/authentication-engine
- run `make install`

## Run
For registering a user: authentication-engine -m 2 -t register -f data/register.json
For authenticating a user: authentication-engine -m 2 -t authenticate -f data/authenticate.json