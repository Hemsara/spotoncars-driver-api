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
	DriverPk         *int
	BookRefNo        *string
	DriverContact    *string
	JobStatus        *int
	BookPickupDtTime *time.Time
	BookPassengerNm  *string
}

func GetBookingsHistory(c *gin.Context) {
	db := initializers.DB

	startDate := time.Now().Truncate(24 * time.Hour)
	endDate := startDate.Add(24 * time.Hour).Add(-time.Nanosecond)
	limit := 40
	var bookings []bookingHistoryDetails

	query := fmt.Sprintf(`
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
