package services

import (
	"net/http"

	"github.com/GabrielEdwinSP/go-simulate-project.git/internal/database"
	"github.com/gin-gonic/gin"
)

type Request struct {
	GajiPokok float32 `json:"gajipokok"`
	Tunjangan float32 `json:"tunjangan"`
}

type Response struct {
	TotalGross   float32 `json:"totalgross"`
	JHT          float32 `json:"jht"`
	JP           float32 `json:"jp"`
	BiayaJabatan float32 `json:"biayajabatan"`
	TotalNet     float32 `json:"totalnet"`
	GajiSetahun  float32 `json:"gajisetahun"`
	TotalSetahun float32 `json:"totalSetahun"`
	PPHperTahun  float32 `json:"pphpertahun"`
	PPHperbulan  float32 `json:"pphperbulan"`
}

func TotalSementara(gaji float32, tunjangan float32) (total_sementara float32, jht float32, jp float32, jabatan float32, totalnet float32) {

	total_sementara = gaji + tunjangan
	jht = gaji * 0.02
	jp = gaji * 0.01
	jabatan = total_sementara * 0.05939
	totalnet = total_sementara - (jht + jp + jabatan)
	return total_sementara, jht, jp, jabatan, totalnet
}

func TotalSetahun(totalnet float32) (setahun float32) {
	setahun = totalnet * 12
	return setahun
}

func PphPerTahun(totalsetahun float32) (pphsetahun float32, pphperbulan float32) {
	switch {
	case totalsetahun <= 60000000:
		pphsetahun = totalsetahun * 0.05
	case totalsetahun > 60000000 && totalsetahun <= 250000000:
		pphsetahun = totalsetahun * 0.15
	case totalsetahun > 250000000 && totalsetahun <= 500000000:
		pphsetahun = totalsetahun * 0.25
	case totalsetahun > 500000000:
		pphsetahun = totalsetahun * 0.3
	}
	pphperbulan = pphsetahun / 12

	return pphsetahun, pphperbulan
}

func PostLogService(c *gin.Context) {
	var request Request
	var response Response

	c.Bind((&request))

	result := database.DB.Create(&request)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//Calculate
	total_sementara, jht2, jp1, jabatan, totalnets := TotalSementara(request.GajiPokok, request.Tunjangan)
	setahun := TotalSetahun(totalnets)
	totalsetahun := setahun - 54000000
	pphsetahun, pphbulanan := PphPerTahun(totalsetahun)

	response.TotalGross = total_sementara
	response.JHT = jht2
	response.JP = jp1
	response.BiayaJabatan = jabatan
	response.TotalNet = totalnets
	response.GajiSetahun = setahun
	response.TotalSetahun = totalsetahun
	response.PPHperTahun = pphsetahun
	response.PPHperbulan = pphbulanan

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})

}
