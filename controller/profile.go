package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"playbox/middleware"
	"playbox/model"
	"playbox/utils"
)

func Profile(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	// get user profile
	r.GET("/farmer-profile", middleware.Authorization(), func(c *gin.Context) {
		strType, _ := c.Get("type")
		ID, _ := c.Get("id")

		if strType == "farmer" {
			var company model.AquaFarmer
			if err := db.Where("id = ?", ID).First(&company).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "Farmer profile", company)
			return
		}
		utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
	})
}
