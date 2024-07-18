package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"spotoncars_server/initializers"
	"spotoncars_server/internal"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminUser struct {
	AdminName     *string
	AdminPassword *string
	userPk        *string
}

func LoginAdmin(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	db := initializers.DB

	var adminUser AdminUser

	query := `
	SELECT AdminName ,AdminPassword ,userPk
	FROM dbo.tbl_AdminTrackingLogin
	WHERE AdminName = @AdminEmail;
	`
	err := db.QueryRow(query, sql.Named("AdminEmail", req.Email)).Scan(&adminUser.AdminName, &adminUser.AdminPassword, &adminUser.userPk)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not find account"})
		return
	}

	if adminUser.AdminName == nil || adminUser.AdminPassword == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not find account"})
		return
	}
	isValid, err := internal.VerifyPassword(req.Password, *adminUser.AdminPassword)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	_, accessToken, err := internal.CreateToken(*adminUser.userPk, *adminUser.AdminName, 24*15*time.Hour)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
