package common

import "github.com/gin-gonic/gin"

func CreateErrorResposeBody(err error) interface{} {
	if err == nil {
		return nil
	}

	return gin.H{"message": err.Error()}
}
