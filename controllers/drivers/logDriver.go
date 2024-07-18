package controllers

import (
	"database/sql"
	"net/http"
	"spotoncars_server/initializers"
	"time"

	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	DvrPk      string `json:"dvrPk"`
	AppVersion string `json:"appVersion"`
}

func LogDriver(c *gin.Context) {
	var req LogRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	db := initializers.DB
	query := `
		UPDATE dbo.Tbl_DriverDetails
		SET AppVersion = @AppVersion, LastLogin = @LastLogin
		WHERE DvrPk = @DvrPk;

		SELECT LastLogin, AppVersion
		FROM dbo.Tbl_DriverDetails
		WHERE DvrPk = @DvrPk;
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database preparation error"})
		return
	}
	defer stmt.Close()

	lastLogin := time.Now().Format(time.RFC3339)
	rows, err := stmt.Query(sql.Named("DvrPk", req.DvrPk), sql.Named("AppVersion", req.AppVersion), sql.Named("LastLogin", lastLogin))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query execution error"})
		return
	}
	defer rows.Close()

	var res struct {
		LastLogin  string `json:"lastLogin"`
		AppVersion string `json:"appVersion"`
	}

	if rows.Next() {
		if err := rows.Scan(&res.LastLogin, &res.AppVersion); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Response scanning error"})
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
	}
}
