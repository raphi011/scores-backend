package auth

import (
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// PasswordCredentials is a email + password combination.
type PasswordCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type credentials struct {
	ClientID    string `json:"client_id"`
	CientSecret string `json:"client_secret"`
}

// OauthUser is returned by the google /userinfo api call.
type OauthUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

// GoogleOAuthConfig reads the google oauth config from the specified file.
// Host is the url the api response should be redirected to.
func GoogleOAuthConfig(configPath, host string) (*oauth2.Config, error) {
	var credentials credentials
	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	redirectURL := host + "/api/auth"

	json.Unmarshal(file, &credentials)
	config := &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.CientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	return config, nil
}
