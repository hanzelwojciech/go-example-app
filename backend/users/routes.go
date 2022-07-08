package users

import (
	"learning-app/common"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ginEngine *gin.Engine) {
	group := ginEngine.Group("/users")

	group.POST(
		"/",
		common.ValidationMiddleware(&UserInviteRequest{}, nil),
		inviteUserRouteHandler,
	)

	group.POST(
		"/activate",
		common.ValidationMiddleware(&UserActivateRequest{}, nil),
		activateUserRouteHandler,
	)

	group.POST(
		"/auth",
		common.ValidationMiddleware(&UserAuthenticateRequest{}, nil),
		authenticateRouteHandler,
	)
}
