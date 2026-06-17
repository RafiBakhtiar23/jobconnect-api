package controllers

import (
	"net/http"

	"jobconnect-api/config"
	"jobconnect-api/models"

	"fmt"
	"jobconnect-api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var user models.User

	// mengambil data JSON dari client
	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data tidak valid",
		})
		return
	}
	fmt.Printf("%+v\n", user)

	// mengubah password menjadi hash
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengenkripsi password",
		})
		return
	}

	// mengganti password asli dengan hash
	user.Password = string(hashedPassword)

	// set role default sebagai user
	user.Role = "user"

	// menyimpan data ke database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat akun",
		})
		return
	}

	// response berhasil
	c.JSON(http.StatusCreated, gin.H{
		"message": "Register berhasil",
		"data":    user,
	})
}

func Login(c *gin.Context) {

	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Input tidak valid",
		})
		return
	}

	var user models.User

	result := config.DB.
		Where("email = ?", input.Email).
		First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau Password salah",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau Password salah",
		})
		return
	}

	fmt.Println("========== LOGIN ==========")
	fmt.Println("ID    :", user.ID)
	fmt.Println("EMAIL :", user.Email)
	fmt.Println("ROLE  :", user.Role)
	fmt.Println("===========================")


	token, err := utils.GenerateToken(
		user.ID,
		user.Role,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   token,
	})
}
