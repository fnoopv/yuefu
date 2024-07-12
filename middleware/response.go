package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type commonResponse struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func Response() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			if ctx.Writer.Status() == 0 {
				ctx.JSON(http.StatusInternalServerError, commonResponse{
					Success: false,
					Message: ctx.Errors.Last().Error(),
				})
				return
			}

			ctx.JSON(ctx.Writer.Status(), commonResponse{
				Success: false,
				Message: ctx.Errors.Last().Error(),
			})
			return
		}

		if ctx.Writer.Status() == 0 {
			ctx.Writer.WriteHeader(http.StatusOK)
		}

		data, _ := ctx.Get("data")

		ctx.JSON(ctx.Writer.Status(), commonResponse{
			Success: true,
			Message: "success",
			Data:    data,
		})
	}
}
