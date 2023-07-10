package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"playbox/middleware"
	"playbox/model"
	"playbox/utils"
)

func Siklus(db *gorm.DB, q *gin.Engine) {
	r := q.Group("api/farmer/:tambak_id/:kolam_id/siklus")

	// get all siklus based on farmerID
	r.GET("/all", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		tambakID := utils.StringToInteger(c.Param("tambak_id"))
		kolamID := utils.StringToInteger(c.Param("kolam_id"))

		var siklus []model.Siklus
		if err := db.Where("aqua_farmer_id = ?", ID).Where("tambak_id = ?", tambakID).Where("kolam_id = ?", kolamID).Find(&siklus).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "all siklus", siklus)

	})

	// create siklus
	r.POST("/create", middleware.Authorization(), func(c *gin.Context) {
		var input model.SiklusInput
		if err := c.BindJSON(&input); err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		ID, _ := c.Get("id")
		tambakID := utils.StringToInteger(c.Param("tambak_id"))
		kolamID := utils.StringToInteger(c.Param("kolam_id"))
		startDate := utils.TimeNowToString()

		newSiklus := model.Siklus{
			AquaFarmerID: ID.(uuid.UUID),
			TambakID:     tambakID,
			KolamID:      kolamID,
			Name:         input.Name,
			StartDate:    startDate,
		}

		if err := db.Create(&newSiklus).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "siklus created", newSiklus)
	})
}
