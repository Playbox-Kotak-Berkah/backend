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
	r := q.Group("/api/farmer/tambak")

	r.GET("/all-tambak", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var tambaks []model.Tambak
		if err := db.Where("aqua_farmer_id = ?", ID).Preload("AquaFarmer").Find(&tambaks).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusFound, "all tambaks", tambaks)
	})

	r.GET("/:tambak_id", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		tambakID := c.Param("tambak_id")

		var tambak model.Tambak
		if err := db.Where("aqua_farmer_id = ?", ID).Where("id = ?", tambakID).Preload("AquaFarmer").First(&tambak).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusFound, "tambak by id", tambak)
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
			CreatedAt:    time.Now(),
		}

		if err := db.Create(&newTambak).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusCreated, "New tambak added", newTambak)
	})
}
