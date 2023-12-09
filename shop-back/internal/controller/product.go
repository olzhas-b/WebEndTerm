package controller

import (
	"github.com/Torebekov/shop-back/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (srv *Server) GetProducts(c *gin.Context) {
	searchText := c.Query("search")
	categoryID, err := strconv.ParseUint(c.Query("categoryID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := srv.core.Product().List(
		searchText,
		categoryID,
		helper.GetUser(c),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
