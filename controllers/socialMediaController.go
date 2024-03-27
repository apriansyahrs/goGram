package controllers

import (
	"goGram/database"
	"goGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	var sosmed models.SocialMedia
	if err := c.ShouldBindJSON(&sosmed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}

	sosmed.UserID = userId

	if err := db.Debug().Create(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               sosmed.ID,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaURL,
		"user_id":          sosmed.UserID,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	sosmedId := c.Param("socialMediaId")
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	var sosmed models.SocialMedia
	if err := db.First(&sosmed, sosmedId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": "Media sosial tidak ditemukan",
		})
		return
	}

	if err := c.ShouldBindJSON(&sosmed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}

	sosmed.UserID = userId

	if err := db.Model(&sosmed).Where("id = ?", sosmedId).Updates(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               sosmed.ID,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaURL,
		"user_id":          sosmed.UserID,
	})
}

func GetAllSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	var sosmeds []models.SocialMedia
	if err := db.Debug().Preload("User").Where("user_id = ?", userId).Find(&sosmeds).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}

	var responseData []models.SocialMediaResponse
	for _, sosmed := range sosmeds {
		socialMedia := models.SocialMediaResponse{
			ID:             sosmed.ID,
			Name:           sosmed.Name,
			SocialMediaURL: sosmed.SocialMediaURL,
			UserID:         sosmed.UserID,
			User: models.UserResponseRelasi{
				ID:       sosmed.User.ID,
				Email:    sosmed.User.Email,
				Username: sosmed.User.Username,
			},
		}
		responseData = append(responseData, socialMedia)
	}

	c.JSON(http.StatusOK, responseData)
}

func GetOneSocialMedia(c *gin.Context) {
	db := database.GetDB()
	sosmedId := c.Param("socialMediaId")
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	var sosmed models.SocialMedia
	if err := db.Debug().Preload("User").Where("user_id = ? AND id = ?", userId, sosmedId).First(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": "Media sosial tidak ditemukan",
		})
		return
	}

	responseData := models.SocialMediaResponse{
		ID:             sosmed.ID,
		Name:           sosmed.Name,
		SocialMediaURL: sosmed.SocialMediaURL,
		UserID:         sosmed.UserID,
		User: models.UserResponseRelasi{
			ID:       sosmed.User.ID,
			Email:    sosmed.User.Email,
			Username: sosmed.User.Username,
		},
	}

	c.JSON(http.StatusOK, responseData)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	sosmedId := c.Param("socialMediaId")
	var sosmed models.SocialMedia
	if err := db.First(&sosmed, sosmedId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": "Media sosial tidak ditemukan",
		})
		return
	}
	db.Delete(&sosmed)
	c.JSON(http.StatusOK, gin.H{
		"message": "Media sosial Anda telah berhasil dihapus",
	})
}
