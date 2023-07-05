package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"playbox/config/database"
	"playbox/controller"
	"playbox/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	databaseConf, err := database.NewDatabase()
	if err != nil {
		panic(err.Error())
	}
	db, err := database.MakeSupaBaseConnectionDatabase(databaseConf)
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
	controller.Profile(db, r)
	controller.Tambak(db, r)

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		panic(err.Error())
	}
}
