package controllers

import (
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"spotoncars_server/initializers"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
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
	isValid, err := verifyPassword(req.Password, *adminUser.AdminPassword)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	_, accessToken, err := CreateToken(*adminUser.userPk, *adminUser.AdminName, 24*15*time.Hour)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func verifyPassword(password, encodedHash string) (bool, error) {

	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, fmt.Errorf("invalid hash format")
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return false, err
	}
	if version != argon2.Version {
		return false, fmt.Errorf("incompatible version of argon2")
	}

	var memory uint32
	var time uint32
	var threads uint8
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, time, memory, threads, uint32(len(hash)))

	return subtle.ConstantTimeCompare(hash, otherHash) == 1, nil
}

func CreateToken(
	driverPK string,
	userName string,
	duration time.Duration,
) (id uuid.UUID, token string, err error) {
	now := time.Now().UTC()

	id, err = uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, "", err
	}

	claims := make(jwt.MapClaims)

	claims["sub"] = id.String()
	claims["exp"] = now.Add(duration).Unix()
	claims["username"] = userName

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("ADMIN_TOKEN_SECRET")))
	if err != nil {
		return uuid.UUID{}, "", err
	}

	return id, token, nil
}
