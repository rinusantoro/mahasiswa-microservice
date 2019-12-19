package main

import (
	"mahasiswa/config"
	"mahasiswa/mahasiswa"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	r := gin.Default()
	db := config.DBInit()
	mahasiswa := mahasiswa.Mahasiswa{DB: db}

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	config.RegisterConsul()
	config.RegisterZipkin()

	r.GET("/mahasiswas", mahasiswa.GetMahasiswas)
	r.POST("/mahasiswas", mahasiswa.CreateMahasiswa)

	r.GET("/healthcheck", config.Healthcheck)

	r.Run() //port 8080
}
