package controllers

import (
	"goGram/database"
	"goGram/helpers"
	"goGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var appJSON = "application/json"

func Register(c *gin.Context) {
	db := database.GetDB()
	contentType := appJSON

	var user models.User
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Permintaan Gagal",
				"message": "Data JSON tidak sesuai dengan format yang diharapkan",
			})
			return
		}
	} else {
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Permintaan Gagal",
				"message": "Data tidak sesuai dengan format yang diharapkan",
			})
			return
		}
	}

	if err := db.Debug().Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}

	response := models.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Age:             user.Age,
		ProfileImageURL: user.ProfileImageURL,
	}

	c.JSON(http.StatusCreated, response)
}

func Login(c *gin.Context) {
	db := database.GetDB()
	contentType := appJSON

	var user models.User
	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	password := user.Password

	if err := db.Debug().Where("email = ?", user.Email).Take(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Tidak Diizinkan",
			"message": "Email atau kata sandi salah",
		})
		return
	}

	comparePass := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Tidak Diizinkan",
			"message": "Email atau kata sandi salah",
		})
		return
	}
	token := helpers.GenerateToken(user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	contentType := appJSON

	var user models.User
	userId := c.MustGet("userId").(uint)
	err := db.First(&user, userId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data tidak ditemukan",
			"message": "Data tidak ditemukan",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	updates := map[string]interface{}{
		"Email":           user.Email,
		"Username":        user.Username,
		"Age":             user.Age,
		"ProfileImageURL": user.ProfileImageURL,
	}
	err = db.Model(&user).Where("id = ?", userId).Updates(updates).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}

	response := models.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Age:             user.Age,
		ProfileImageURL: user.ProfileImageURL,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()

	var user models.User
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	err := db.Delete(&user, userId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Akun Anda berhasil dihapus",
	})
}
