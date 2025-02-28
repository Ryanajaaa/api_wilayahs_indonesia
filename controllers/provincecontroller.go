package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Ryanajaa/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

// Struktur untuk API Response
type ProvinceAPI struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Fetch dari API dan Simpan ke Database
func FetchProvinces(c *gin.Context) {
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	var apiProvinces []ProvinceAPI
	if err := json.NewDecoder(resp.Body).Decode(&apiProvinces); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	// Simpan ke database
	for _, apiProvince := range apiProvinces {
		province := models.Province{
			Code: apiProvince.ID, // Pastikan ID dari API menjadi Code
			Name: apiProvince.Name,
		}
		models.DB.Create(&province)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Data inserted"})
}

// Ambil data dari database
func GetProvinces(c *gin.Context) {
	var provinces []models.Province
	models.DB.Find(&provinces)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully get data",
		"data":    provinces,
	})
}
