package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api"
)

func SendNotification(externalIDs []string, message string) (string, *onesignal.CreateNotificationSuccessResponse, error) {
	notification := *onesignal.NewNotification(os.Getenv("ONE_SIGNAL_APP_ID"))

	// Optionally, include external user IDs if provided

	// notification.SetIncludeExternalUserIds([]string{"344"})
	notification.IncludedSegments = []string{"All"} // Target all users

	englishMessage := message
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
		return "", nil, err
	}

	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNotification`: %v\n", resp)

	return "Notification sent successfully", resp, nil
}
