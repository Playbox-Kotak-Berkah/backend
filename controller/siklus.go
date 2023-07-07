package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"playbox/model"
	"playbox/utils"
)

func Siklus(db *gorm.DB, q *gin.Engine) {
	r := q.Group("api/farmer/:tambak_id/:kolam_id")

	// get siklus per doc
	r.GET("/all-latest", func(c *gin.Context) {
		tambakID := utils.StringToInteger(c.Param("tambak_id"))
		kolamID := utils.StringToInteger(c.Param("kolam_id"))

		var siklus []model.Siklus

		if err := db.Where("tambak_id = ?", tambakID).Where("kolam_id = ?", kolamID).Order("created_at asc").Find(&siklus).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "all siklus", siklus)
	})
}
