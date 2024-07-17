package controllers

import (
	"fmt"
	"net/http"
	"spotoncars_server/initializers"

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

	c.JSON(http.StatusOK, gin.H{
		"drivers": drivers,
	})
}