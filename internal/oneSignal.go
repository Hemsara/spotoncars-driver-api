package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api"
)

func sendNotification() (string, error, *onesignal.CreateNotificationSuccessResponse) {
	notification := *onesignal.NewNotification(os.Getenv("ONE_SIGNAL_APP_ID"))
	// notification.SetExternalId("322")

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
