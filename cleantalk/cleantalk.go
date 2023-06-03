package cleantalk

import (
	gin "github.com/gin-gonic/gin"
)

func Handler(r *gin.Engine) {
	r.POST("/simple", HandleTalk)
}
