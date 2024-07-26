package controllers

import (
	"database/sql"
	"net/http"
	"spotoncars_server/initializers"
	"spotoncars_server/internal"

	"github.com/gin-gonic/gin"
)

type NotificationRequest struct {
	DriverIDs []string `json:"driver_ids"`
	Message   string   `json:"message"`
}

func SendNotification(c *gin.Context) {
	var req NotificationRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	externalID := req.DriverIDs
	message := req.Message

	isSent, _, err := internal.SendNotification(externalID, message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send notification"})
		return
	}
	if isSent {
		db := initializers.DB

		query := `
			INSERT INTO dbo.Tbl_DriverNotifications (DvrPk, NotificationText)
			VALUES (@DriverID, @Message)
			`
		stmt, err := db.Prepare(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare SQL statement"})
			return
		}
		defer stmt.Close()

		for _, driverID := range externalID {
			_, err := stmt.Exec(sql.Named("DriverID", driverID), sql.Named("Message", message))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert notification into database"})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": isSent,
	})

}
