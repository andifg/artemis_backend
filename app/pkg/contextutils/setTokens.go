package contextutils

import (
	"github.com/andifg/artemis_backend/app/pkg/auth"
	"github.com/gin-gonic/gin"
)

func SetTokens(c *gin.Context, tokens auth.OidcTokens) {

	c.SetCookie("access_token", tokens.AccessToken, 2592000, "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, 2592000, "/", "localhost", false, true)
	c.SetCookie("id_token", tokens.IdToken, 2592000, "/", "localhost", false, false)

}
