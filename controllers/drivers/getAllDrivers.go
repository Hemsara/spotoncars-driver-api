package controllers

import (
	"context"
	"fmt"
	"net/http"
	"spotoncars_server/initializers"

	"os"

	"github.com/OneSignal/onesignal-go-api"
	"github.com/gin-gonic/gin"
)

type driverDetails struct {
	DvrFName     *string
	DvrLName     *string
	DvrContactNo *string
	DvrEmailId   *string
	DvrLicNo     *string
	DvrPk        *string
	LastLogin    *string
	AppVersion   *string
}

func GetAllDrivers(c *gin.Context) {
	db := initializers.DB
	var drivers []driverDetails

	query := `

SELECT TOP 20
    DvrFName,
    DvrLName,
    DvrContactNo,
    DvrEmailId,
    DvrLicNo,
    DvrPk,
    LastLogin,
    AppVersion
FROM
    tbl_DriverDetails
WHERE
    DvrDelId = 0;

	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var driver driverDetails

		if err := rows.Scan(&driver.DvrFName, &driver.DvrLName, &driver.DvrEmailId, &driver.DvrContactNo, &driver.DvrLicNo, &driver.DvrPk, &driver.LastLogin, &driver.AppVersion); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan failed"})
			return
		}
		drivers = append(drivers, driver)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Rows iteration failed"})
		return
	}

	successMessage, err, _ := sendNotification()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"drivers":        drivers,
		"successMessage": successMessage,
	})
}

func sendNotification() (string, error, *onesignal.CreateNotificationSuccessResponse) {
	notification := *onesignal.NewNotification(os.Getenv("ONE_SIGNAL_APP_ID"))

	// Optionally, include external user IDs if provided

	// notification.SetIncludeExternalUserIds([]string{"344"})
	notification.IncludedSegments = []string{"All"} // Target all users

	englishMessage := "This is a notification in English!"
	contents := onesignal.StringMap{
		En: &englishMessage,
	}

	notification.SetContents(contents)

	configuration := onesignal.NewConfiguration()

	apiClient := onesignal.NewAPIClient(configuration)

	appAuth := context.WithValue(context.Background(), onesignal.AppAuth, os.Getenv("ONE_SIGNAL_APP_KEY"))

	resp, r, err := apiClient.DefaultApi.CreateNotification(appAuth).Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNotification`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err, nil
	}

	// Print response from `CreateNotification`
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNotification`: %v\n", resp)

	// Return success message, error, and response
	return "Notification sent successfully", nil, resp
}
