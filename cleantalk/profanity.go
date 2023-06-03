package cleantalk

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func HandleTalk(c *gin.Context) {
	text := c.PostForm("text")

	response := gin.H{
		"message": "The text does not contain obscene words",
		"obscene": false,
	}

	if isProfanity(text) {
		response["message"] = "The text contains an obscene word"
		response["obscene"] = true
		c.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
