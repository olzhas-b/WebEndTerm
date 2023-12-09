package controller

import (
	"github.com/Torebekov/shop-back/internal/helper"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (srv *Server) GetCart(c *gin.Context) {
	results, err := srv.core.Cart().List(helper.GetUser(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (srv *Server) AddCart(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = srv.core.Cart().Add(models.Cart{ProductID: productID, UserID: helper.GetUser(c)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (srv *Server) RemoveCart(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = srv.core.Cart().Remove(models.Cart{ProductID: productID, UserID: helper.GetUser(c)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
