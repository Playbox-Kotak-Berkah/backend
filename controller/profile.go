package controller

import (
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"playbox/middleware"
	"playbox/model"
	"playbox/utils"
	"time"
)

func Profile(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	SupaBaseClient := supabasestorageuploader.NewSupabaseClient(
		os.Getenv("SUPABASE_PROJECT_URL"),
		os.Getenv("SUPABASE_PROJECT_API_KEY"),
		os.Getenv("SUPABASE_PROJECT_STORAGE_NAME"),
		os.Getenv("SUPABASE_STORAGE_FOLDER"),
	)

	// get farmer profile
	r.GET("/farmer-profile", middleware.Authorization(), func(c *gin.Context) {
		strType, _ := c.Get("type")
		ID, _ := c.Get("id")

		if strType == "farmer" {
			var farmer model.AquaFarmer
			if err := db.Where("id = ?", ID).First(&farmer).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "Farmer profile", farmer)
			return
		}

		if strType == "user" {
			var user model.User
			if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "User profile", user)
			return

		}
		utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
	})

	//r.GET("/user-profile", middleware.Authorization(), func(c *gin.Context) {
	//	strType, _ := c.Get("type")
	//	ID, _ := c.Get("id")
	//
	//	if strType == "user" {
	//		var user model.User
	//		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
	//			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
	//			return
	//		}
	//
	//		utils.HttpRespSuccess(c, http.StatusOK, "User profile", user)
	//		return
	//	}
	//	utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
	//})

	r.PATCH("/farmer-profile", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		strType, _ := c.Get("type")

		if strType == "user" {
			var input model.UserUpdateProfileInput
			if err := c.BindJSON(&input); err != nil {
				utils.HttpRespFailed(c, http.StatusBadRequest, err.Error())
				return
			}

			var user model.User
			if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			user.Name = input.Name
			user.UpdatedAt = time.Now()

			if err := db.Save(&user).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "User profile updated", user)
			return
		}

		if strType == "farmer" {
			var input model.AquaFarmerEditProfileInput
			if err := c.BindJSON(&input); err != nil {
				utils.HttpRespFailed(c, http.StatusBadRequest, err.Error())
				return
			}

			var farmer model.AquaFarmer
			if err := db.Where("id = ?", ID).First(&farmer).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			farmer.Name = input.Name
			farmer.UpdatedAt = time.Now()

			if err := db.Save(&farmer).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "Farmer profile updated", farmer)
			return

		}
		utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
	})

	r.PATCH("/farmer-profile-picture", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		strType, _ := c.Get("type")

		if strType == "user" {
			var user model.User
			if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			photo, _ := c.FormFile("photo")
			uploaded, err := SupaBaseClient.Upload(photo)
			if err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			user.Picture = uploaded

			if err := db.Save(&user).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "User profile updated", user)
			return
		}

		if strType == "farmer" {
			var farmer model.AquaFarmer
			if err := db.Where("id = ?", ID).First(&farmer).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			photo, _ := c.FormFile("photo")
			uploaded, err := SupaBaseClient.Upload(photo)
			if err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			farmer.Picture = uploaded

			if err := db.Save(&farmer).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "Farmer profile updated", farmer)
			return
		}
		utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
	})
}
