package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/osmanakol/oauth2-golang/pkg/google"
)

type GoogleHandler interface {
	GetAuthCodeUrl(c *fiber.Ctx) error
	GetToken(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

type GoogleImpl struct {
	google *google.GoogleAuth2
}

func NewGoogleHandler(auth *google.GoogleAuth2) *GoogleImpl {
	o := &GoogleImpl{
		google: auth,
	}

	return o
}

func (o *GoogleImpl) GetAuthCodeUrl(c *fiber.Ctx) error {
	authUrl := o.google.GetLoginUrl()

	return c.Status(200).JSON(authUrl)
}

func (o *GoogleImpl) GetToken(c *fiber.Ctx) error {
	authCode := c.FormValue("code")
	fmt.Printf("", c.Queries())
	token, err := o.google.AuthorizationCodeFlow(authCode)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	userInfo, err := o.google.GetUserInformation(token.AccessToken)

	if err != nil {
		return c.Status(400).JSON(err)
	}
	fmt.Printf("%s", userInfo)
	return c.Status(200).JSON(userInfo)
}

func (o *GoogleImpl) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Query("refreshToken")
	token, err := o.google.NewAccessTokenFromRefreshToken(refreshToken)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(token)
}
