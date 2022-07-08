package users

import (
	"errors"
	"learning-app/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func inviteUserRouteHandler(c *gin.Context) {
	body := c.MustGet("requestBody").(*UserInviteRequest)

	_, err := InviteUser(*body)

	errResponse := common.CreateErrorResposeBody(err)

	switch {
	case err == nil:
		c.Status(http.StatusOK)
		return
	case errors.Is(err, ErrEmailExists):
		c.JSON(http.StatusConflict, errResponse)
		return
	default:
		c.JSON(http.StatusBadRequest, errResponse)
	}
}

func activateUserRouteHandler(c *gin.Context) {
	body := c.MustGet("requestBody").(*UserActivateRequest)

	err := ActivateUser(*body)
	errResponse := common.CreateErrorResposeBody(err)

	switch {
	case err == nil:
		c.Status(200)
		return
	case errors.Is(err, ErrTokenNotFound):
		c.JSON(http.StatusNotFound, errResponse)
		return
	default:
		c.JSON(http.StatusBadRequest, errResponse)
	}

}

func authenticateRouteHandler(c *gin.Context) {
	body := c.MustGet("requestBody").(*UserAuthenticateRequest)

	result, err := Authenticate(*body)
	errResponse := common.CreateErrorResposeBody(err)

	switch {
	case err == nil:
		c.JSON(http.StatusOK, result)
		return
	case errors.Is(err, ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, errResponse)
		return
	default:
		c.JSON(http.StatusBadRequest, errResponse)
	}
}
