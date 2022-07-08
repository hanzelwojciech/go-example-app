package common

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bindSchema(name string, schema interface{}, c *gin.Context) {
	var err error

	switch name {
	case "requestBody":
		err = c.ShouldBind(&schema)
		fmt.Println("Accessed request body mapper")
	case "requestQuery":
		err = c.ShouldBindQuery(&schema)
	default:
		err = errors.New("cannot resolve proper schema type")
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Set(name, schema)
}

func ValidationMiddleware(
	bodySchema interface{},
	querySchema interface{},
) gin.HandlerFunc {
	return func(c *gin.Context) {
		if bodySchema != nil {
			bindSchema("requestBody", bodySchema, c)
		}

		if querySchema != nil {
			bindSchema("requestQuery", querySchema, c)
		}
	}
}
