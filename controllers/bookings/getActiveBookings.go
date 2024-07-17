package controllers

import (
	"fmt"
	"net/http"
	"spotoncars_server/initializers"
	"time"

	"github.com/gin-gonic/gin"
)

type bookingDetails struct {
	DriverName       *string
	DriverPk         *int
	BookRefNo        *string
	DriverContact    *string
	JobStatus        *int
	BookPickupDtTime *time.Time
}

func GetActiveBookings(c *gin.Context) {
	db := initializers.DB
	var bookings []bookingDetails

	startOfToday := time.Now().Truncate(24 * time.Hour)
	endOfToday := startOfToday.Add(24 * time.Hour).Add(-time.Nanosecond)

	query := `
SELECT 
    DriverName,
    DriverPk,
    BookRefNo,
    DriverContact,
    JobStatus,
    BookPickupDtTime
FROM
    Tbl_BookingDetails
WHERE
    JobStatus IN (2, 3, 4)
	AND BookPickupDtTime >= '` + startOfToday.Format("2006-01-02 15:04:05") + `'
    AND BookPickupDtTime <= '` + endOfToday.Format("2006-01-02 15:04:05") + `'
ORDER BY
    BookPickupDtTime DESC
OFFSET 0 ROWS
FETCH NEXT 20 ROWS ONLY;


	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var booking bookingDetails

		if err := rows.Scan(&booking.DriverName, &booking.DriverPk, &booking.BookRefNo, &booking.DriverContact, &booking.JobStatus, &booking.BookPickupDtTime); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan failed"})
			return
		}
		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Rows iteration failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bookings": bookings,
	})
}
