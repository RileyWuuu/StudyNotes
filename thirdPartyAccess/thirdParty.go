package thirdpartyaccess

import (
	"encoding/base64"
	"math/rand"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	config = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
)

// Create a random state for authentication
func RandomState() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, config.RedirectURL, http.StatusTemporaryRedirect)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

}

// type googleUser struct {
// 	Sub           string `json:"sub"`
// 	Name          string `json:"name"`
// 	GivenName     string `json:"given_name"`
// 	FamilyName    string `json:"family_name"`
// 	Profile       string `json:"profile"`
// 	Picture       string `json:"picture"`
// 	Email         string `json:"email"`
// 	EmailVerified bool   `json:"email_verified"`
// 	Gender        string `json:"gender"`
// 	Hd            string `json:"hd"`
// }

//	func getOauthURL() (*oauth2.Config, string) {
//		options := CreateClientOptions("google", "http://localhost:8080/callback")
//		config = &oauth2.Config{
//			ClientID:     options.getID(),
//			ClientSecret: options.getSecret(),
//			RedirectURL:  options.getRedirectURL(),
//			Scopes: []string{
//				"https://www.googleapis.com/auth/userinfo.email",
//				"https://www.googleapis.com/auth/userinfo.profile",
//			},
//			Endpoint: google.Endpoint,
//		}
//	}
func GenerateAccessToken(w http.ResponseWriter, r *http.Request) {

}
