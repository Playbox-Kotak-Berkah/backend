package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
	"net/http"
	"os"
	"playbox/middleware"
	"playbox/model"
	"playbox/utils"
)

func Paket(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/farmer/paket")

	r.GET("/all-paket", middleware.Authorization(), func(c *gin.Context) {
		strType, _ := c.Get("type")

		if strType != "farmer" {
			utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
			return
		}

		var paktes []model.Paket
		if err := db.Find(&paktes).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "all paket", paktes)
	})

	r.POST("/:paket_id/payment", middleware.Authorization(), func(c *gin.Context) {
		strType, _ := c.Get("type")

		if strType != "farmer" {
			utils.HttpRespFailed(c, http.StatusUnauthorized, "Not authorized")
			return
		}

		paketID := c.Param("paket_id")

		var paket model.Paket
		if err := db.Where("id = ?", paketID).First(&paket).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		midtransClient := coreapi.Client{}
		midtransClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
		orderID := utils.RandomOrderID()

		req := &coreapi.ChargeReq{
			PaymentType: "gopay",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderID,
				GrossAmt: paket.Price,
			},
			Gopay: &coreapi.GopayDetails{
				EnableCallback: true,
				CallbackUrl:    "https://example.com/callback",
			},
		}

		resp, err := midtransClient.ChargeTransaction(req)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		ID, _ := c.Get("id")
		var farmer model.AquaFarmer
		if err := db.Where("id = ?", ID).First(&farmer).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		farmer.IsVerified = true

		if err := db.Save(&farmer).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Payment success", resp)
	})
}
