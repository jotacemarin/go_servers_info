package services

import (
	"database/sql"

	"../db"
	"../models"
)

// InsertDomain : func
func InsertDomain(domain models.Domain) (models.Domain, error) {
	var newdomain models.Domain
	database := db.Db
	queryDomain, errdbP := database.Prepare("INSERT INTO domain(servers_changed, ssl_grade, previus_ssl_grade, logo, title, is_down) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if errdbP != nil {
		return newdomain, errdbP
	}
	var lastID int
	errqEx := queryDomain.QueryRow(domain.ServerChanged, domain.SslGrade, domain.PreviusSslGrade, domain.Logo, domain.Title, domain.IsDown).Scan(&lastID)
	if errqEx != nil {
		return newdomain, errqEx
	}
	defer queryDomain.Close()
	for _, server := range domain.Servers {
		_, errIS := InsertServer(server, lastID)
		if errIS != nil {
			return newdomain, errIS
		}
	}
	return newdomain, nil
}

// GetLast : func
func GetLast(domain models.Domain) (models.Domain, error) {
	var newdomain models.Domain
	database := db.Db
	row := database.QueryRow("SELECT servers_changed, ssl_grade, previus_ssl_grade, logo, title, is_down FROM domain WHERE lower(title) LIKE lower('%' || $1 || '%') ORDER BY id DESC LIMIT 1;", domain.Title)
	errScan := row.Scan(&newdomain.ServerChanged, &newdomain.SslGrade, &newdomain.PreviusSslGrade, &newdomain.Logo, &newdomain.Title, &newdomain.IsDown)
	switch errScan {
	case sql.ErrNoRows:
		return newdomain, nil
	case nil:
		return newdomain, nil
	default:
		return newdomain, errScan
	}
}

// HistoryDomains : func
func HistoryDomains() ([]models.Domain, error) {
	var domains []models.Domain
	database := db.Db
	queryResult, errQ := database.Query("SELECT servers_changed, ssl_grade, previus_ssl_grade, logo, title, is_down FROM domain WHERE id IN (SELECT MAX(id) as lastId FROM domain GROUP BY title);")
	if errQ != nil {
		return domains, errQ
	}
	defer queryResult.Close()
	for queryResult.Next() {
		var domain models.Domain
		errSc := queryResult.Scan(&domain.ServerChanged, &domain.SslGrade, &domain.PreviusSslGrade, &domain.Logo, &domain.Title, &domain.IsDown)
		if errSc != nil {
			return domains, errSc
		}
		domains = append(domains, domain)
	}
	return domains, nil
}
