package auth

import (
	"learning-app/common"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var response = common.CreateErrorResposeBody(ErrUnauthorized)

func AuthMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.Replace(header, "Bearer ", "", 1)

		claims, err := VerifyDecodeToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

		if len(roles) > 0 {
			containsRole := checkIfClaimsContainsRole(roles, claims)
			if !containsRole {
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			}
		}

	}
}

func checkIfClaimsContainsRole(roles []string, claims *CustomTokenClaims) bool {
	claimsRole := getRoleFromClaims(claims)
	result := false

	for _, role := range roles {
		if role == claimsRole {
			result = true
		}
	}

	return result
}

func getRoleFromClaims(claims *CustomTokenClaims) string {
	payload := reflect.ValueOf(claims.Payload).Elem()
	role := payload.FieldByName("Role").Interface().(string)
	return role
}
