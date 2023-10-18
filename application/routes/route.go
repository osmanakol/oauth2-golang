package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/osmanakol/oauth2-golang/application/handler"
)

func Register(r *fiber.App) {
	v1 := r.Group("/v1")

	h, err := handler.NewHandler(
		handler.WithGoogleRoute(),
	)

	if err != nil {
		panic(err)
	}

	v1.Get("/auth/google", h.GoogleRoute.GetAuthCodeUrl)
	v1.Get("/auth/google/refresh", h.GoogleRoute.RefreshToken)
	v1.Get("/auth/google/callback", h.GoogleRoute.GetToken)
}
