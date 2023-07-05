package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"playbox/middleware"
	"playbox/model"
	"playbox/utils"
	"time"
)

func Tambak(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/farmer")

	r.GET("/all-tambak", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var tambaks []model.Tambak
		if err := db.Where("aqua_farmer_id = ?", ID).Find(&tambaks).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "all tambaks", tambaks)
	})

	r.POST("/add-tambak", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		var input model.AddTambak
		if err := c.BindJSON(&input); err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		newTambak := model.Tambak{
			Name:         input.Name,
			AquaFarmerID: ID.(uuid.UUID),
			CreatedAt:    time.Time{},
		}

		if err := db.Create(&newTambak).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "New tambak added", newTambak)
	})
}
