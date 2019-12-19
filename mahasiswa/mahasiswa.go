package mahasiswa

import (
	"mahasiswa/model"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.opencensus.io/trace"
)

type Mahasiswa struct {
	DB *gorm.DB
}

type mahasiswa struct {
	NIM         string `json:"nim"`
	Nama        string `json:"nama"`
	TempatLahir string `json:"tempat_lahir"`
	Agama       string `json:"agama"`
}

func (p *Mahasiswa) GetMahasiswas(c *gin.Context) {
	db := p.DB
	var mahasiswas []model.Mahasiswa

	db.Find(&mahasiswas)

	c.JSON(200, gin.H{
		"data": mahasiswas,
	})

	serviceb(c)
}

func (p *Mahasiswa) CreateMahasiswa(c *gin.Context) {
	var request model.Mahasiswa

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	p.DB.Create(&request)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func serviceb(c *gin.Context) {
	_, span := trace.StartSpan(c, "/mahasiswas")
	defer span.End()
	time.Sleep(time.Duration(rand.Intn(800)+200) * time.Millisecond)
}
