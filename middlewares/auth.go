package middlewares

import (
	"goGram/database"
	"goGram/helpers"
	"goGram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifiyToken, err := helpers.VerifiyToken(c)
		_ = verifiyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Akses Terbatas",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifiyToken)
		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Permintaan Gagal",
				"message": "Parameter tidak valid",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data tidak ditemukan",
				"message": "Data tidak ditemukan",
			})
			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Akses Terbatas",
				"message": "Anda tidak diizinkan mengakses data ini",
			})
			return
		}
		c.Next()

	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Permintaan Gagal",
				"message": "Parameter tidak valid",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data tidak ditemukan",
				"message": "Data tidak ditemukan",
			})
			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Akses Terbatas",
				"message": "Anda tidak diizinkan mengakses data ini",
			})
			return
		}
		c.Next()

	}
}
func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialmediaId, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Permintaan Gagal",
				"message": "Parameter tidak valid",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(socialmediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data tidak ditemukan",
				"message": "Data tidak ditemukan",
			})
			return
		}

		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Akses Terbatas",
				"message": "Anda tidak diizinkan mengakses data ini",
			})
			return
		}
		c.Next()

	}
}
