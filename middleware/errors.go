package middleware

import (
	"net/http"

	"github.com/ctirouzh/tiny-url/pkg/apperr"
	"github.com/gin-gonic/gin"
)

func ErrorsMiddleware(errorType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var appError *apperr.AppError
		errors := c.Errors.ByType(errorType)
		if len(errors) > 0 {
			err := errors[0].Err
			switch err := err.(type) {
			case *apperr.AppError:
				appError = err
			default:
				appError = &apperr.AppError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			c.IndentedJSON(appError.Code, appError)
			c.Abort()
			return
		}
	}
}
