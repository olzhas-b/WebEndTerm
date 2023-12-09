package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (srv *Server) GetCategories(c *gin.Context) {
	results, err := srv.core.Category().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
