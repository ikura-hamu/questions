package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	traqoauth2 "github.com/ras0q/traq-oauth2"
)

type OAuthHandler interface {
	AuthorizeHandler(c echo.Context) error
	CallbackHandler(c echo.Context) error
}

var (
	traqClientId    = os.Getenv("TRAQ_CLIENT_ID")
	traqRedirectUrl = os.Getenv("TRAQ_REDIRECT_URL")
	conf            = traqoauth2.NewConfig(traqClientId, traqRedirectUrl)
)

func AuthorizeHandler(c echo.Context) error {
	codeVerifier, err := traqoauth2.GenerateCodeVerifier()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to generate code verifier: %v", err))
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get seeeion: %v", err))
	}

	sess.Values["code_verifier"] = codeVerifier
	sess.Save(c.Request(), c.Response())

	codeChallengeMethod := traqoauth2.CodeChallengeMethod(c.QueryParam("method"))
	if codeChallengeMethod == "" {
		codeChallengeMethod = traqoauth2.CodeChallengePlain
	}

	codeChallenge, err := traqoauth2.GenerateCodeChallenge(codeVerifier, codeChallengeMethod)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to generate code challenge: %v", err))
	}

	authCodeURL := conf.AuthCodeURL(
		c.QueryParam("state"),
		traqoauth2.WithCodeChallenge(codeChallenge),
		traqoauth2.WithCodeChallengeMethod(codeChallengeMethod),
	)

	return c.Redirect(http.StatusSeeOther, authCodeURL)
}

func CallbackHandler(c echo.Context) error {
	sess, err := session.Get("sessions", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get session: %v", err))
	}

	codeVerifier, ok := sess.Values["code_verifier"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	code := c.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "no code")
	}

	ctx := c.Request().Context()
	token, err := conf.Exchange(ctx, code, traqoauth2.WithCodeVerifier(codeVerifier))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to exchange code into token: %v", err))
	}

	sess.Values["access_token"] = token

	sess.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, "ok")
}
