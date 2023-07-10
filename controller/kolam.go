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

func Kolam(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/farmer/:tambak_id")

	r.GET("/all-kolam", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		tambakID := utils.StringToInteger(c.Param("tambak_id"))

		var kolams []model.Kolam
		if err := db.Where("aqua_farmer_id = ?", ID).Where("tambak_id = ?", tambakID).Preload("AquaFarmer").Find(&kolams).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "all kolams", kolams)
	})

	r.GET("/:kolam_id", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		tambakID := c.Param("tambak_id")
		kolamID := c.Param("kolam_id")

		var kolam model.Kolam
		if err := db.Where("aqua_farmer_id = ?", ID).Where("tambak_id = ?", tambakID).Where("id = ?", kolamID).Preload("AquaFarmer").First(&kolam).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "kolam by id", kolam)
	})

	r.POST("/add-kolam", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		tambakID := utils.StringToInteger(c.Param("tambak_id"))

		var input model.AddKolam
		if err := c.BindJSON(&input); err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		newKolam := model.Kolam{
			Name:              input.Name,
			AquaFarmerID:      ID.(uuid.UUID),
			TambakID:          tambakID,
			LampuTambakStatus: false,
			KincirAirStatus:   false,
			KeranAirStatus:    false,
			CreatedAt:         time.Now(),
		}

		if err := db.Create(&newKolam).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusCreated, "New kolam added", newKolam)
	})

	// change control status
	r.POST("/:kolam_id", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		tambakID := utils.StringToInteger(c.Param("tambak_id"))
		kolamID := c.Param("kolam_id")

		var kolam model.Kolam
		if err := db.Where("aqua_farmer_id = ?", ID).Where("tambak_id = ?", tambakID).Where("id = ?", kolamID).Preload("AquaFarmer").First(&kolam).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		water := utils.StringToBool(c.Query("water"))
		bulb := utils.StringToBool(c.Query("bulb"))
		fan := utils.StringToBool(c.Query("fan"))

		kolam.KeranAirStatus = water
		kolam.KincirAirStatus = fan
		kolam.LampuTambakStatus = bulb

		if err := db.Save(&kolam).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "kolam updated successfully", kolam)
	})

}
