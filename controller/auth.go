package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"os"
	"playbox/model"
	"playbox/utils"
	"time"
)

func FarmerRegister(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	r.POST("/farmer-register", func(c *gin.Context) {
		var input model.AquaFarmerRegisterInput

		if err := c.BindJSON(&input); err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		var existingFarmer model.AquaFarmer
		if err := db.Where("email = ?", input.Email).First(&existingFarmer).Error; err == nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Email already registered")
			return
		}

		if err := db.Where("phone = ?", input.Phone).First(&existingFarmer).Error; err == nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Phone number already registered")
			return
		}

		var existingUser model.User
		if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Email already registered")
			return
		}

		if err := db.Where("phone = ?", input.Phone).First(&existingUser).Error; err == nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Phone number already registered")
			return
		}

		if input.Name == "" {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Name cannot be empty")
			return
		}

		if input.Password != input.ConfirmPassword {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, "Password and confirm password does not match")
			return
		}

		hashedPassword, err := utils.Hash(input.Password)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
		}

		newFarmer := model.AquaFarmer{
			ID:        uuid.New(),
			Name:      input.Name,
			Phone:     input.Phone,
			Email:     input.Email,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
		}

		if err := db.Create(&newFarmer).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusCreated, "Success", input)
	})
}

func FarmerLogin(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	r.POST("/farmer-login", func(c *gin.Context) {
		var input model.AquaFarmerLoginInput

		if err := c.BindJSON(&input); err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		var user model.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			// utils.HttpRespFailed(c, http.StatusNotFound, "User not found")
		}

		var farmer model.AquaFarmer
		if err := db.Where("email = ?", input.Email).First(&farmer).Error; err != nil {
			// utils.HttpRespFailed(c, http.StatusNotFound, "Email is not registered")
			// return
		}

		var accountType string

		if user.ID != uuid.Nil && utils.CompareHash(input.Password, user.Password) {
			accountType = "user"
			token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"id":   user.ID,
				"type": accountType,
				"exp":  time.Now().Add(time.Hour).Unix(),
			})

			strToken, err := token.SignedString([]byte(os.Getenv("TOKEN")))
			if err != nil {
				utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "Parsed token", gin.H{
				"name":  user.Name,
				"token": strToken,
				"type":  accountType,
			})
			return

		}

		if farmer.ID != uuid.Nil && utils.CompareHash(input.Password, farmer.Password) {
			accountType = "farmer"
			token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"id":   farmer.ID,
				"type": accountType,
				"exp":  time.Now().Add(time.Hour).Unix(),
			})

			strToken, err := token.SignedString([]byte(os.Getenv("TOKEN")))
			if err != nil {
				utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "Parsed token", gin.H{
				"name":  farmer.Name,
				"token": strToken,
				"type":  accountType,
			})
			return

		}

		utils.HttpRespFailed(c, http.StatusForbidden, "Wrong email or password")
		return

		// original code
		//var existingFarmer model.AquaFarmer
		//if err := db.Where("email = ?", input.Email).First(&existingFarmer).Error; err != nil {
		//	utils.HttpRespFailed(c, http.StatusBadRequest, "Email does not exist")
		//	return
		//}
		//
		//if err := utils.CompareHash(input.Password, existingFarmer.Password); err != true {
		//	utils.HttpRespFailed(c, http.StatusBadRequest, "Password does not match")
		//	return
		//}
		//
		//accountType := "farmer"
		//token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		//	"id":   existingFarmer.ID,
		//	"type": accountType,
		//	"exp":  time.Now().Add(time.Hour).Unix(),
		//})
		//
		//strToken, err := token.SignedString([]byte(os.Getenv("TOKEN")))
		//if err != nil {
		//	utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
		//	return
		//}
		//
		//utils.HttpRespSuccess(c, http.StatusOK, "Parsed token", gin.H{
		//	"name":  existingFarmer.Name,
		//	"token": strToken,
		//	"type":  accountType,
		//})
	})
}

func UserRegister(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	r.POST("/user-register", func(c *gin.Context) {
		var input model.UserRegisterInput
		if err := c.BindJSON(&input); err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		var existingFarmer model.AquaFarmer
		if err := db.Where("email = ?", input.Email).First(&existingFarmer).Error; err == nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Email already registered")
			return
		}

		if err := db.Where("phone = ?", input.Phone).First(&existingFarmer).Error; err == nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Phone number already registered")
			return
		}

		var existingUser model.User
		if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Email already registered")
			return
		}

		if err := db.Where("phone = ?", input.Phone).First(&existingUser).Error; err == nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Phone number already registered")
			return
		}

		if input.Name == "" {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Name cannot be empty")
			return
		}

		if input.Password != input.ConfirmPassword {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, "Password and confirm password does not match")
			return
		}

		hashedPassword, err := utils.Hash(input.Password)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
		}

		newUser := model.User{
			ID:        uuid.New(),
			Name:      input.Name,
			Phone:     input.Phone,
			Email:     input.Email,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
		}

		if err := db.Create(&newUser).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusCreated, "New user created", newUser)
	})
}

func UserLogin(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	r.POST("/user-login", func(c *gin.Context) {
		var input model.UserLoginInput

		if err := c.BindJSON(&input); err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		var existingUser model.User
		if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Email does not exist")
			return
		}

		if err := utils.CompareHash(input.Password, existingUser.Password); err != true {
			utils.HttpRespFailed(c, http.StatusBadRequest, "Password does not match")
			return
		}

		accountType := "user"
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"id":   existingUser.ID,
			"type": accountType,
			"exp":  time.Now().Add(time.Hour).Unix(),
		})

		strToken, err := token.SignedString([]byte(os.Getenv("TOKEN")))
		if err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Parsed token", gin.H{
			"name":  existingUser.Name,
			"token": strToken,
			"type":  accountType,
		})
	})
}
