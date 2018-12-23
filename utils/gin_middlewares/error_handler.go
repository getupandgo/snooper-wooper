package gin_middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Error struct {
	Status  int
	Message string
}

func AppErrorReporter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		controllerError := c.Errors.Last()

		if controllerError != nil {
			var formattedError Error

			switch controllerError.Type {
			case gin.ErrorTypePublic:
				log.Info().
					Err(controllerError).
					Msg("")

				formattedError = Error{http.StatusBadRequest, controllerError.Error()}

			case gin.ErrorTypeBind:
				log.Info().
					Err(controllerError).
					Msg("")

				formattedError = Error{http.StatusBadRequest, "Invalid param"}

			case gin.ErrorTypePrivate:
				log.Error().
					Err(controllerError).
					Msg("")

				formattedError = Error{http.StatusInternalServerError, "Server error"}

			default:
				log.Error().
					Err(controllerError).
					Msg("Unhandled error")

				formattedError = Error{http.StatusInternalServerError, "Server error"}
			}

			c.AbortWithStatusJSON(formattedError.Status, formattedError.Message)
		}
	}
}
