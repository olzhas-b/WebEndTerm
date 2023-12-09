package controller

import (
	"github.com/Torebekov/shop-back/internal/helper"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (srv *Server) GetProducts(c *gin.Context) {
	searchText := c.Query("search")
	categoryID, _ := strconv.ParseUint(c.Query("categoryID"), 10, 64)

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

func (srv *Server) GetProduct(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := srv.core.Product().ByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (srv *Server) AddProduct(c *gin.Context) {
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = srv.core.Product().Add(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (srv *Server) RemoveProduct(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = srv.core.Product().Remove(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
