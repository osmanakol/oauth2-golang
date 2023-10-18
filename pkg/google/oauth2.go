package google

import (
	"context"
	"encoding/json"
	"fmt"
	config "github.com/osmanakol/oauth2-golang/configuration"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
)

type GoogleAuth2 struct {
	conf  *oauth2.Config
	state string
}

func NewGoogleOauth2() GoogleAuth2 {
	googleConfiguration := config.Env.GetGoogleConfig()
	return GoogleAuth2{
		conf: &oauth2.Config{
			ClientID:     googleConfiguration.CID,
			ClientSecret: googleConfiguration.CSECRET,
			RedirectURL:  googleConfiguration.REDIRECT_URL,
			Scopes:       []string{"email", "profile"},
			Endpoint:     google.Endpoint,
		},
		state: "",
	}
}

func (googleAuth *GoogleAuth2) GetLoginUrl() string {
	return googleAuth.conf.AuthCodeURL(googleAuth.state, oauth2.AccessTypeOffline, oauth2.ApprovalForce, oauth2.S256ChallengeOption("123"))
}

func (googleAuth *GoogleAuth2) AuthorizationCodeFlow(code string) (*oauth2.Token, error) {
	token, err := googleAuth.conf.Exchange(context.Background(), code, oauth2.VerifierOption("123"))

	if err != nil {
		fmt.Printf("err %s", err.Error())
		return &oauth2.Token{}, err
	}

	return token, err
}

func (googleAuth *GoogleAuth2) NewAccessTokenFromRefreshToken(refreshToken string) (*oauth2.Token, error) {
	token := oauth2.Token{
		RefreshToken: refreshToken,
	}

	newToken := googleAuth.conf.TokenSource(oauth2.NoContext, &token)

	return newToken.Token()
}

type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func (googleAuth *GoogleAuth2) GetUserInformation(accessToken string) (UserInfo, error) {
	client := googleAuth.conf.Client(context.Background(), &oauth2.Token{
		AccessToken: accessToken,
	})
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", accessToken)

	resp, err := client.Get(rootUrl)
	defer resp.Body.Close()
	if err != nil {
		return UserInfo{}, err
	}

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}, err
	}

	var userInfo UserInfo
	err = json.Unmarshal(contents, &userInfo)

	if err != nil {
		return UserInfo{}, err
	}
	return userInfo, nil
}
