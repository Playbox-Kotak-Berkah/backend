package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"playbox/config"
	"playbox/controller"
	"playbox/middleware"
)

func main() {
	databaseConf, err := config.NewDatabase()
	if err != nil {
		panic(err.Error())
	}
	db, err := config.MakeSupaBaseConnectionDatabase(databaseConf)
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()

	r.Use(middleware.CORS())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"env": os.Getenv("ENV"),
		})
	})

	controller.FarmerRegister(db, r)
	controller.FarmerLogin(db, r)

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		panic(err.Error())
	}
}
