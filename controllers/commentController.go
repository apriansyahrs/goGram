package controllers

import (
	"goGram/database"
	"goGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := appJSON

	Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userId

	if err := db.Debug().Create(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":       Comment.ID,
		"message":  Comment.Message,
		"photo_id": Comment.PhotoID,
		"user_id":  Comment.UserID,
	})

}

func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	comments := []models.Comment{}

	if err := db.Debug().Preload("User").Preload("Photo").Where("user_id = ?", userId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}
	var response []gin.H
	for _, comment := range comments {
		commentData := gin.H{
			"id":       comment.ID,
			"message":  comment.Message,
			"photo_id": comment.PhotoID,
			"user_id":  comment.UserID,
		}
		userData := gin.H{
			"id":       comment.User.ID,
			"email":    comment.User.Email,
			"username": comment.User.Username,
		}
		photoData := gin.H{
			"id":        comment.Photo.ID,
			"title":     comment.Photo.Title,
			"caption":   comment.Photo.Caption,
			"photo_url": comment.Photo.PhotoURL,
			"user_id":   comment.Photo.UserID,
		}
		commentData["user"] = userData
		commentData["photo"] = photoData
		response = append(response, commentData)
	}
	c.JSON(http.StatusOK, response)
}

func GetOneComment(c *gin.Context) {
	db := database.GetDB()
	commentId := c.Param("commentId")
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	comment := models.Comment{}

	if err := db.Debug().Preload("User").Preload("Photo").Where("id = ? AND user_id = ?", commentId, userId).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Tidak Ditemukan",
			"message": "Komentar tidak ditemukan atau Anda tidak memiliki izin untuk mengaksesnya",
		})
		return
	}

	commentData := gin.H{
		"id":       comment.ID,
		"message":  comment.Message,
		"photo_id": comment.PhotoID,
		"user_id":  comment.UserID,
	}
	userData := gin.H{
		"id":       comment.User.ID,
		"email":    comment.User.Email,
		"username": comment.User.Username,
	}
	photoData := gin.H{
		"id":        comment.Photo.ID,
		"title":     comment.Photo.Title,
		"caption":   comment.Photo.Caption,
		"photo_url": comment.Photo.PhotoURL,
		"user_id":   comment.Photo.UserID,
	}
	commentData["user"] = userData
	commentData["photo"] = photoData

	c.JSON(http.StatusOK, commentData)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()

	commentId := c.Param("commentId")

	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := appJSON

	Comment := models.Comment{}
	db.First(&Comment, commentId)

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userId

	if err := db.Model(&Comment).Where("id = ?", commentId).Updates(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Permintaan Gagal",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":       Comment.ID,
		"message":  Comment.Message,
		"photo_id": Comment.PhotoID,
		"user_id":  Comment.UserID,
	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentId := c.Param("commentId")
	Comment := models.Comment{}
	db.First(&Comment, commentId)
	db.Delete(&Comment)
	c.JSON(http.StatusOK, gin.H{
		"message": "Komentar Anda berhasil dihapus",
	})
}
