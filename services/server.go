package services

import (
	"fmt"
	"strconv"

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

// GetServers : func
func GetServers(domainID int64) ([]models.Server, error) {
	dID := strconv.FormatInt(domainID, 10)
	var servers []models.Server
	database := db.Db
	queryResult, errQ := database.Query("SELECT domain, address, ssl_grade, country, owner FROM server WHERE domain = $1;", dID)
	if errQ != nil {
		return servers, fmt.Errorf("GetServer 0: %s", errQ)
	}
	defer queryResult.Close()
	for queryResult.Next() {
		var server models.Server
		errSc := queryResult.Scan(&server.Domain, &server.Address, &server.SslGrade, &server.Country, &server.Owner)
		if errSc != nil {
			return servers, fmt.Errorf("GetServers 1: %s", errSc)
		}
		servers = append(servers, server)
	}
	return servers, nil
}
