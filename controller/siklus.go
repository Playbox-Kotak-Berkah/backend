package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"playbox/middleware"
	"playbox/model"
	"playbox/utils"
	"time"
)

func Siklus(db *gorm.DB, q *gin.Engine) {
	r := q.Group("api/farmer/:tambak_id/:kolam_id/")

	// get all siklus based on farmerID
	r.GET("/all-siklus", middleware.Authorization(), func(c *gin.Context) {
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
	r.POST("/add-siklus", middleware.Authorization(), func(c *gin.Context) {
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

	// input data
	r.POST(":siklus_id/input-data", middleware.Authorization(), func(c *gin.Context) {
		var input model.InputSiklusHarian
		if err := c.BindJSON(&input); err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		ID, _ := c.Get("id")
		tambakID := utils.StringToInteger(c.Param("tambak_id"))
		kolamID := utils.StringToInteger(c.Param("kolam_id"))

		siklusID := utils.StringToInteger(c.Param("siklus_id"))

		log.Println(ID)
		log.Println(tambakID)
		log.Println(kolamID)
		log.Println(siklusID)

		newSiklusHarian := model.SiklusHarian{
			SiklusID:      siklusID,
			Tanggal:       utils.TimeNowToString(),
			PHRealtime:    utils.CalculateRealtime(input.PHPagi, input.PHSiang, input.PHMalam),
			PHPagi:        input.PHPagi,
			PHSiang:       input.PHSiang,
			PHMalam:       input.PHMalam,
			SuhuRealtime:  utils.CalculateRealtime(input.SuhuPagi, input.SuhuSiang, input.SuhuMalam),
			SuhuPagi:      input.SuhuPagi,
			SuhuSiang:     input.SuhuSiang,
			SuhuMalam:     input.SuhuMalam,
			DORealtime:    utils.CalculateRealtime(input.DOPagi, input.DOSiang, input.DOMalam),
			DOPagi:        input.DOPagi,
			DOSiang:       input.DOSiang,
			DOMalam:       input.DOMalam,
			GaramRealtime: utils.CalculateRealtime(input.GaramPagi, input.GaramSiang, input.GaramMalam),
			GaramPagi:     input.GaramPagi,
			GaramSiang:    input.GaramSiang,
			GaramMalam:    input.GaramMalam,
			CreatedAt:     time.Now(),
		}

		if err := db.Create(&newSiklusHarian).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "siklus harian created", newSiklusHarian)
	})
}
