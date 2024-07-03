package controllers

import (
	"fmt"
	"net/http"
	"spotoncars_server/initializers"

	"github.com/gin-gonic/gin"
)

type bookingDetails struct {
	BookRefNo       *string `json:"bookRefNo"`
	BookPassengerNm *string `json:"bookPassengerNm"`
}

func GetActiveBookings(c *gin.Context) {
	db := initializers.DB
	var bookings []bookingDetails

	query := `
		SELECT
    BookRefNo,
    BookPassengerNm
FROM
    Tbl_BookingDetails
ORDER BY
    BookRefNo
OFFSET 0 ROWS FETCH NEXT 10 ROWS ONLY;

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
		if err := rows.Scan(&booking.BookRefNo, &booking.BookPassengerNm); err != nil {
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
