package controllers

import (
	"fmt"
	"net/http"
	"spotoncars_server/initializers"
	"time"

	"github.com/gin-gonic/gin"
)

type bookingHistoryDetails struct {
	DriverName       *string
	DvrPk            *int
	BookRefNo        *string
	DriverContact    *string
	JobStatus        *int
	BookPickupDtTime *time.Time
	BookPassengerNm  *string
}

func GetBookingsHistory(c *gin.Context) {
	db := initializers.DB
	filterBy := c.Query("filterBy")
	bookRef := c.Query("bookRef")

	startDate, endDate := getDateFilters(filterBy)

	limit := 40
	var bookings []bookingHistoryDetails

	var query string

	if bookRef != "" {

		query = fmt.Sprintf(`
				SELECT
					DriverName,
					BookRefNo,
					DriverContact,
					JobStatus,
					BookPickupDtTime,
					BookPassengerNm
				FROM
					Tbl_BookingDetails
				WHERE
					JobStatus = 5 AND
					BookRefNo ='%s'

				ORDER BY
					BookPickupDtTime DESC;
			`, bookRef)

	} else {

		query = fmt.Sprintf(`
		SELECT TOP %d
			DriverName,
			BookRefNo,
			DriverContact,
			JobStatus,
			BookPickupDtTime,
			BookPassengerNm
		FROM
			Tbl_BookingDetails
		WHERE
			JobStatus = 5
			AND BookPickupDtTime >= '%s'
			AND BookPickupDtTime <= '%s'
		ORDER BY
			BookPickupDtTime DESC;
	`, limit, startDate.Format("2006-01-02T15:04:05"), endDate.Format("2006-01-02T15:04:05"))
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var booking bookingHistoryDetails

		if err := rows.Scan(&booking.DriverName, &booking.BookRefNo, &booking.DriverContact, &booking.JobStatus, &booking.BookPickupDtTime, &booking.BookPassengerNm); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan failed"})
			return
		}

		bookings = append(bookings, booking)
	}

	c.JSON(http.StatusOK, gin.H{
		"bookings": bookings,
	})
}

func getDateFilters(filterBy string) (startDateFilter, endDateFilter time.Time) {
	now := time.Now()

	switch filterBy {
	case "today":
		startDateFilter = now.Truncate(24 * time.Hour)
		endDateFilter = startDateFilter.Add(24*time.Hour - time.Nanosecond)

	case "yesterday":
		startDateFilter = now.AddDate(0, 0, -1).Truncate(24 * time.Hour)
		endDateFilter = startDateFilter.Add(24*time.Hour - time.Nanosecond)

	case "7 days":
		startDateFilter = now.AddDate(0, 0, -7).Truncate(24 * time.Hour)
		endDateFilter = now.Truncate(24 * time.Hour).Add(24*time.Hour - time.Nanosecond)

	case "one month":
		startDateFilter = now.AddDate(0, -1, 0).Truncate(24 * time.Hour)
		endDateFilter = now.Truncate(24 * time.Hour).Add(24*time.Hour - time.Nanosecond)

	default:
		startDateFilter = now.Truncate(24 * time.Hour)
		endDateFilter = startDateFilter.Add(24*time.Hour - time.Nanosecond)
	}

	return startDateFilter, endDateFilter
}
