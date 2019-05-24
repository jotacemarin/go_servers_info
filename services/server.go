package services

import (
	"../db"
	"../models"
)

// InsertServer : func
func InsertServer(server models.Server, domainID int) (models.Server, error) {
	var newserver models.Server
	database := db.Db
	query, errdbP := database.Prepare("INSERT INTO server(domain, address, ssl_grade, country, owner) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if errdbP != nil {
		return newserver, errdbP
	}
	_, errqEx := query.Exec(domainID, server.Address, server.SslGrade, server.Country, server.Owner)
	if errqEx != nil {
		return newserver, errqEx
	}
	defer query.Close()
	return newserver, nil
}
