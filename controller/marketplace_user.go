package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"playbox/middleware"
	"playbox/model"
	"playbox/utils"
	"strconv"
)

func MarketPlaceUser(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/marketplace-user")
	r.GET("/all", middleware.Authorization(), func(c *gin.Context) {
		page := c.Query("page")   // Get the page number from query parameter
		limit := c.Query("limit") // Get the limit from query parameter

		// Set default values if page and limit are not provided
		if page == "" {
			page = "1" // Default to the first page
		}

		if limit == "" {
			limit = "10" // Default to 10 products per page
		}

		// Convert page and limit to integers
		pageNum, err := strconv.Atoi(page)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Invalid page number")
			return
		}

		limitNum, err := strconv.Atoi(limit)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Invalid limit value")
			return
		}

		var products []model.Product

		// Calculate the offset based on the page and limit
		offset := (pageNum - 1) * limitNum

		// Retrieve the products with pagination using offset and limit
		if err := db.Offset(offset).Limit(limitNum).Find(&products).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Products", products)
	})
}
