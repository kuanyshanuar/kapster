package handlers


import(
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type statusResponse struct{ 
	Status string `json:"status"`
}

type errorResponse struct{ 
		Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, msg string, statusCode int ) {
				logrus.Error(msg)
				c.AbortWithStatusJSON(statusCode, errorResponse{msg})

}