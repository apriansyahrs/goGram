package controllers

import (
	"goGram/database"
	"goGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreatePhoto(c *gin.Context) {

	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := appJSON

	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId

	if err := db.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        Photo.ID,
		"caption":   Photo.Caption,
		"title":     Photo.Title,
		"photo_url": Photo.PhotoURL,
		"user_id":   Photo.UserID,
	})

}

func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	photos := []models.Photo{}

	if err := db.Debug().Preload("User").Where("user_id = ?", userId).Find(&photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}
	var responseData []models.PhotoResponse
	for _, photo := range photos {
		photoData := models.PhotoResponse{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
			UserID:   photo.UserID,
			User: models.UserResponseRelasi{
				ID:       photo.User.ID,
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		}
		responseData = append(responseData, photoData)
	}

	c.JSON(http.StatusOK, responseData)
}

func GetOnePhoto(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	photoId := c.Param("photoId")
	photo := models.Photo{}

	if err := db.Debug().Preload("User").Where("user_id = ? AND id = ?", userId, photoId).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": "Foto tidak ditemukan",
		})
		return
	}

	responseData := models.PhotoResponse{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
		User: models.UserResponseRelasi{
			ID:       photo.User.ID,
			Email:    photo.User.Email,
			Username: photo.User.Username,
		},
	}

	c.JSON(http.StatusOK, responseData)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()

	photoId := c.Param("photoId")

	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := appJSON

	Photo := models.Photo{}
	db.First(&Photo, photoId)

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId

	if err := db.Model(&Photo).Where("id = ?", photoId).Updates(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":        Photo.ID,
		"caption":   Photo.Caption,
		"title":     Photo.Title,
		"photo_url": Photo.PhotoURL,
		"user_id":   Photo.UserID,
	})

}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photoId := c.Param("photoId")
	Photo := models.Photo{}
	db.First(&Photo, photoId)
	db.Delete(&Photo)
	c.JSON(http.StatusOK, gin.H{
		"message": "Foto Anda berhasil dihapus",
	})
}
