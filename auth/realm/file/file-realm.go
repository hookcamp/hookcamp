package file

import (
	"errors"
	"fmt"

	"github.com/frain-dev/convoy/auth"
	"github.com/frain-dev/convoy/config"
)

var (
	ErrCredentialNotFound = errors.New("credential not found")
)

type BasicAuth struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     auth.Role `json:"role"`
}

type APIKeyAuth struct {
	APIKey string    `json:"api_key"`
	Role   auth.Role `json:"role"`
}

type FileRealm struct {
	Name   string       `json:"name"`
	Basic  []BasicAuth  `json:"basic"`
	APIKey []APIKeyAuth `json:"api_key"`
}

func (r *FileRealm) GetName() string {
	return r.Name
}

func (r *FileRealm) Authenticate(cred *auth.Credential) (*auth.AuthenticatedUser, error) {
	switch cred.Type {
	case auth.CredentialTypeBasic:
		for _, b := range r.Basic {
			if cred.Username != b.Username {
				continue
			}

			if cred.Password != b.Password {
				continue
			}

			authUser := &auth.AuthenticatedUser{
				AuthenticatedByRealm: r.Name,
				Credential:           *cred,
				Role:                 b.Role,
			}
			return authUser, nil
		}
		return nil, ErrCredentialNotFound
	case auth.CredentialTypeAPIKey:
		for _, b := range r.APIKey {
			if cred.APIKey != b.APIKey {
				continue
			}

			authUser := &auth.AuthenticatedUser{
				AuthenticatedByRealm: r.Name,
				Credential:           *cred,
				Role:                 b.Role,
			}
			return authUser, nil
		}
		return nil, ErrCredentialNotFound
	default:
		return nil, fmt.Errorf("unsupported credential type: %s", cred.Type.String())
	}
}

// NewFileRealm constructs a new File Realm authenticator
func NewFileRealm(opts *config.FileRealmOption) (*FileRealm, error) {
	fr := &FileRealm{}
	for _, basicAuth := range opts.Basic {
		fr.Basic = append(fr.Basic, BasicAuth{
			Username: basicAuth.Username,
			Password: basicAuth.Password,
			Role: auth.Role{
				Type:  basicAuth.Role.Type,
				Group: basicAuth.Role.Group,
			},
		})
	}

	for _, basicAuth := range opts.APIKey {
		fr.APIKey = append(fr.APIKey, APIKeyAuth{
			APIKey: basicAuth.APIKey,
			Role: auth.Role{
				Type:  basicAuth.Role.Type,
				Group: basicAuth.Role.Group,
			},
		})
	}

	if len(fr.Basic) == 0 && len(fr.APIKey) == 0 {
		return nil, fmt.Errorf("no authentication data supplied in file realm '%s", fr.Name)
	}

	return fr, nil
}
