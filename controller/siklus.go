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
		if err := db.Where("aqua_farmer_id = ?", ID).Where("tambak_id = ?", tambakID).Where("kolam_id = ?", kolamID).Where("end_date IS NULL OR end_date = ?", "").Find(&siklus).Error; err != nil {
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

		newSiklus := model.Siklus{
			AquaFarmerID: ID.(uuid.UUID),
			TambakID:     tambakID,
			KolamID:      kolamID,
			Name:         input.Name,
			StartDate:    input.StartDate,
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
		siklusID := utils.StringToInteger(c.Param("siklus_id"))

		var siklus model.Siklus
		if err := db.Where("id = ?", siklusID).Where("aqua_farmer_id = ?", ID).Find(&siklus).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		var siklusHarian model.SiklusHarian
		if err := db.Where("siklus_id = ?", siklusID).Order("created_at desc").First(&siklusHarian).Error; err != nil {
			//utils.HttpRespFailed(c, http.StatusFound, err.Error())
			//return
		}

		timeNow := utils.TimeNowToString()
		missingDates, days := utils.MissingDates(siklusHarian.Tanggal, timeNow)

		for _, date := range missingDates {
			newSiklusHarian := model.SiklusHarian{
				SiklusID:      siklusID,
				Tanggal:       date,
				PHRealtime:    0,
				PHPagi:        0,
				PHSiang:       0,
				PHMalam:       0,
				SuhuRealtime:  0,
				SuhuPagi:      0,
				SuhuSiang:     0,
				SuhuMalam:     0,
				DORealtime:    0,
				DOPagi:        0,
				DOSiang:       0,
				DOMalam:       0,
				GaramRealtime: 0,
				GaramPagi:     0,
				GaramSiang:    0,
				GaramMalam:    0,
				CreatedAt:     time.Now(),
			}

			log.Println(days)

			if err := db.Create(&newSiklusHarian).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
				return
			}
		}

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

		if siklusHarian.Tanggal == timeNow {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, "Today's data already inputted")
			return
		}

		if err := db.Create(&newSiklusHarian).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "siklus harian created", newSiklusHarian)
	})

	// get all siklus harian based on siklus id
	r.GET(":siklus_id/all-siklus-harian", middleware.Authorization(), func(c *gin.Context) {
		siklusID := utils.StringToInteger(c.Param("siklus_id"))

		var siklusHarian []model.SiklusHarian
		if err := db.Where("siklus_id = ?", siklusID).Find(&siklusHarian).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "all siklus harian", siklusHarian)
	})

	r.GET(":siklus_id/latest", middleware.Authorization(), func(c *gin.Context) {
		siklusID := utils.StringToInteger(c.Param("siklus_id"))

		var siklus model.Siklus
		if err := db.Where("id = ?", siklusID).First(&siklus).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusFound, err.Error())
			return
		}

		var siklusHarian model.SiklusHarian
		if err := db.Where("siklus_id = ?", siklusID).Order("created_at desc").First(&siklusHarian).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusFound, err.Error())
			return
		}

		siklusHarianResponse := model.SiklusHarianResponse{
			ID:              siklusHarian.ID,
			SiklusID:        siklusHarian.SiklusID,
			Tanggal:         siklusHarian.Tanggal,
			PHRealtime:      siklusHarian.PHRealtime,
			PHPagi:          siklusHarian.PHPagi,
			PHSiang:         siklusHarian.PHSiang,
			PHMalam:         siklusHarian.PHMalam,
			SuhuRealtime:    siklusHarian.SuhuRealtime,
			SuhuPagi:        siklusHarian.SuhuPagi,
			SuhuSiang:       siklusHarian.SuhuSiang,
			SuhuMalam:       siklusHarian.SuhuMalam,
			DORealtime:      siklusHarian.DORealtime,
			DOPagi:          siklusHarian.DOPagi,
			DOSiang:         siklusHarian.DOSiang,
			DOMalam:         siklusHarian.DOMalam,
			GaramRealtime:   siklusHarian.GaramRealtime,
			GaramPagi:       siklusHarian.GaramPagi,
			GaramSiang:      siklusHarian.GaramSiang,
			GaramMalam:      siklusHarian.GaramMalam,
			SiklusStartDate: siklus.StartDate,
			DOC:             utils.CountDays(siklus.StartDate, utils.TimeNowToString()),
			CreatedAt:       siklusHarian.CreatedAt,
			UpdatedAt:       siklusHarian.UpdatedAt,
		}

		utils.HttpRespSuccess(c, http.StatusOK, "latest siklus harian", siklusHarianResponse)
	})

	// end siklus
	r.POST(":siklus_id/akhiri-siklus", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		tambakID := utils.StringToInteger(c.Param("tambak_id"))
		kolamID := utils.StringToInteger(c.Param("kolam_id"))
		siklusID := utils.StringToInteger(c.Param("siklus_id"))

		var siklus model.Siklus
		if err := db.Where("aqua_farmer_id = ?", ID).Where("tambak_id = ?", tambakID).Where("kolam_id = ?", kolamID).Where("id = ?", siklusID).First(&siklus).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		siklus.EndDate = utils.TimeNowToString()

		if err := db.Save(&siklus).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "siklus ended", siklus)
	})
}
