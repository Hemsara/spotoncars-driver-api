package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"spotoncars_server/initializers"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminUser struct {
	AdminName     *string
	AdminPassword *string
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
	SELECT AdminName ,AdminPassword
	FROM dbo.tbl_AdminTrackingLogin
	WHERE AdminName = @AdminEmail;
	`
	err := db.QueryRow(query, sql.Named("AdminEmail", req.Email)).Scan(&adminUser.AdminName, &adminUser.AdminPassword)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not find account"})
		return
	}

	if adminUser.AdminName == nil || adminUser.AdminPassword == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not find account"})
		return
	}

	accessToken := "sample_access_token"

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
