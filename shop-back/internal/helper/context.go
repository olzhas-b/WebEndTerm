package helper

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) uint64 {
	return c.GetUint64("userID")
}
