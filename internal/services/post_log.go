package services

import (
	"net/http"

	"github.com/GabrielEdwinSP/go-simulate-project.git/internal/database"
	"github.com/GabrielEdwinSP/go-simulate-project.git/internal/dto"
	"github.com/gin-gonic/gin"
)

func TotalSementara(gaji float32, tunjangan float32) *dto.CalculateStruct {
	var calculate dto.CalculateStruct
	calculate = dto.CalculateStruct{
		Total_sementara: gaji + tunjangan,
		JHT:             gaji * 0.02,
		JP:              gaji * 0.01,
		Jabatan:         calculate.Total_sementara * 0.05939,
		Totalnet:        calculate.Total_sementara - (calculate.JHT + calculate.JP + calculate.Jabatan),
	}
	// total_sementara =
	// jht = gaji * 0.02
	// jp = gaji * 0.01
	// jabatan = total_sementara * 0.05939
	// totalnet = total_sementara - (jht + jp + jabatan)
	// return total_sementara, jht, jp, jabatan, totalnet
	return &calculate
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
	var request dto.Request

	c.Bind((&request))

	result := database.DB.Create(&request)
	if result.Error != nil {
		c.Status(400)
		return
	}

	calculate := dto.CalculateStruct{}
	//Calculate
	calculate = *TotalSementara(request.GajiPokok, request.Tunjangan)
	setahun := TotalSetahun(calculate.Totalnet)
	totalsetahun := setahun - 54000000
	pphsetahun, pphbulanan := PphPerTahun(totalsetahun)

	response := dto.Response{
		TotalGross:   calculate.Total_sementara,
		JHT:          calculate.Total_sementara,
		JP:           calculate.JP,
		BiayaJabatan: calculate.Jabatan,
		TotalNet:     calculate.Totalnet,
		GajiSetahun:  setahun,
		TotalSetahun: totalsetahun,
		PPHperTahun:  pphsetahun,
		PPHperbulan:  pphbulanan,
	}
	// response.TotalGross = calculate.Total_sementara
	// response.JHT = calculate.JHT
	// response.JP = calculate.JP
	// response.BiayaJabatan = calculate.Jabatan
	// response.TotalNet = calculate.Totalnet
	// response.GajiSetahun = setahun
	// response.TotalSetahun = totalsetahun
	// response.PPHperTahun = pphsetahun
	// response.PPHperbulan = pphbulanan

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})

	resultResponse := database.DB.Create(&response)
	if resultResponse.Error != nil {
		c.Status(400)
		return
	}

}
