package controller

import (
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"os"
	"playbox/middleware"
	"playbox/model"
	"playbox/utils"
)

func MarketPlaceFarmer(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/market-farmer/")

	SupaBaseClient := supabasestorageuploader.NewSupabaseClient(
		os.Getenv("SUPABASE_PROJECT_URL"),
		os.Getenv("SUPABASE_PROJECT_API_KEY"),
		os.Getenv("SUPABASE_PROJECT_STORAGE_NAME"),
		os.Getenv("SUPABASE_STORAGE_FOLDER"),
	)

	r.GET("/all", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		strType, _ := c.Get("type")

		if strType != "farmer" {
			utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
			return
		}

		var products []model.Product
		if err := db.Where("aqua_farmer_id = ?", ID).Find(&products).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "all products", products)
	})

	r.GET("/products/:id", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		strType, _ := c.Get("type")

		if strType != "farmer" {
			utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
			return
		}

		productID := c.Param("id")

		var product model.Product
		if err := db.Where("id = ?", productID).Where("aqua_farmer_id = ?", ID).First(&product).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "product", product)
	})

	r.POST("/add", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		strType, _ := c.Get("type")

		if strType != "farmer" {
			utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
			return
		}

		name := c.PostForm("name")
		photo, _ := c.FormFile("photo")
		price := utils.StringToFloat64(c.PostForm("price"))
		description := c.PostForm("description")
		sold := utils.GenerateRandomInt(10, 200)
		rating := utils.GenerateRandomFloat(0.1, 5.0)

		linkPhoto, _ := SupaBaseClient.Upload(photo)

		newProduct := model.Product{
			AquaFarmerID: ID.(uuid.UUID),
			Name:         name,
			Photo:        linkPhoto,
			Price:        price,
			Description:  description,
			Sold:         sold,
			Rating:       rating,
		}

		if err := db.Create(&newProduct).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "product added", newProduct)
	})

	r.PATCH("/update/:id", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		strType, _ := c.Get("type")

		if strType != "farmer" {
			utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
			return
		}

		productID := c.Param("id")

		name := c.PostForm("name")
		photo, _ := c.FormFile("photo")
		price := utils.StringToFloat64(c.PostForm("price"))
		description := c.PostForm("description")

		linkPhoto, _ := SupaBaseClient.Upload(photo)

		var product model.Product
		if err := db.Where("id = ?", productID).Where("aqua_farmer_id = ?", ID).First(&product).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		product.Name = name
		product.Photo = linkPhoto
		product.Price = price
		product.Description = description

		if err := db.Save(&product).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "product updated", product)
	})

	r.DELETE("/delete/:id", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")
		strType, _ := c.Get("type")

		if strType != "farmer" {
			utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
			return
		}

		productID := c.Param("id")

		var product model.Product
		if err := db.Where("id = ?", productID).Where("aqua_farmer_id = ?", ID).First(&product).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		if err := db.Delete(&product).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "product deleted", product)
	})
}
