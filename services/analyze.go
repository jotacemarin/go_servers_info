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
	query, errdbP := database.Prepare("INSERT INTO domain(servers_changed, ssl_grade, previus_ssl_grade, logo, title, is_down) VALUES ($1, $2, $3, $4, $5, $6)")
	if errdbP != nil {
		return newdomain, errdbP
	}
	_, errqEx := query.Exec(domain.ServerChanged, domain.SslGrade, domain.PreviusSslGrade, domain.Logo, domain.Title, domain.IsDown)
	if errqEx != nil {
		return newdomain, errqEx
	}
	defer query.Close()
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
