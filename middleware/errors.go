package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func ErrorsMiddleware(errorType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var appError *AppError
		errors := c.Errors.ByType(errorType)
		if len(errors) > 0 {
			err := errors[0].Err
			switch err := err.(type) {
			case *AppError:
				appError = err
			default:
				appError = &AppError{
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
