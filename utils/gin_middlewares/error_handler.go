package gin_middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Status  int
	Message string
}

func AppErrorReporter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		controllerError := c.Errors.Last()

		var formattedError Error
		switch controllerError.Type {
		case gin.ErrorTypePublic:
			formattedError = Error{http.StatusBadRequest, controllerError.Error()}
		case gin.ErrorTypeBind:
			formattedError = Error{http.StatusBadRequest, "Invalid param"}
		case gin.ErrorTypePrivate:
			formattedError = Error{http.StatusInternalServerError, "Invalid param"}
		default:

		}

		c.AbortWithStatusJSON(formattedError.Status, formattedError.Message)
	}
}
