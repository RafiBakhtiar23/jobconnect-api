package controllers

import (
	"fmt"
	"net/http"

	"jobconnect-api/config"
	"jobconnect-api/models"

	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {

	var job models.Job

	if err := c.ShouldBindJSON(&job); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data tidak valid",
		})

		return
	}

	if err := config.DB.Create(&job).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat lowongan",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Lowongan berhasil dibuat",
		"data":    job,
	})
}

func GetJobs(c *gin.Context) {

	var jobs []models.Job

	// mengambil query dari URL
	search := c.Query("search")

	// mengambil page dan limit
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")

	var pageInt int
	var limitInt int

	fmt.Sscanf(page, "%d", &pageInt)
	fmt.Sscanf(limit, "%d", &limitInt)

	offset := (pageInt - 1) * limitInt
	fmt.Println("PAGE:", pageInt)
	fmt.Println("LIMIT:", limitInt)
	fmt.Println("OFFSET:", offset)

	query := config.DB

	// filter pencarian
	if search != "" {
		query = query.Where(
			"judul LIKE ? OR company LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
	}

	// pagination
	query.
		Limit(limitInt).
		Offset(offset).
		Find(&jobs)

	c.JSON(http.StatusOK, gin.H{
		"message": "Data lowongan berhasil diambil",
		"page":    pageInt,
		"limit":   limitInt,
		"data":    jobs,
	})
}

func GetJobByID(c *gin.Context) {
	var job models.Job

	id := c.Param("id")

	result := config.DB.First(&job, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Lowongan tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": job,
	})
}

func UpdateJob(c *gin.Context) {
	var job models.Job

	id := c.Param("id")

	result := config.DB.First(&job, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Lowongan tidak ditemukan",
		})
		return
	}

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data tidak valid",
		})
		return
	}

	config.DB.Save(&job)

	c.JSON(http.StatusOK, gin.H{
		"message": "Lowongan berhasil diperbarui",
		"data":    job,
	})
}

func DeleteJob(c *gin.Context) {
	var job models.Job

	id := c.Param("id")

	result := config.DB.First(&job, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Lowongan tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&job)

	c.JSON(http.StatusOK, gin.H{
		"message": "Lowongan berhasil dihapus",
	})
}
