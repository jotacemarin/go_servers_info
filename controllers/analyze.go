package analyzecontroller

import (
	"log"
	"net/http"

	"../commons"
	"../models"
	"../services"
)

// Index : funct
func Index(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("host")
	if len(queryParam) == 0 {
		commons.BuilderJSON(w, false, 0, nil)
	}
	er, errDI := commons.GetDomainInfo(queryParam)
	if errDI != nil {
		commons.BuilderJSON(w, false, 0, nil)
	}
	var domain models.Domain
	if title, logo, errP := commons.GetPageData(queryParam); len(title) > 0 || len(logo) > 0 {
		if errP != nil {
			log.Print(errP)
			commons.BuilderJSON(w, false, 0, nil)
		}
		domain.Title = title
		domain.Logo = logo
	}
	if endpointsLength := len(er.Endpoints); endpointsLength > 0 {
		servers, errW := commons.GetWhois(er.Endpoints)
		if errW != nil {
			commons.BuilderJSON(w, false, 0, nil)
		}
		domain.SslGrade = commons.GetPoorSslGrade(servers)
		domain.Servers = servers
		domain.IsDown = false
	} else {
		domain.IsDown = true
	}
	getLastDomain, errGL := services.GetLast(domain)
	if errGL != nil {
		commons.BuilderJSON(w, false, 0, nil)
	}
	domain.PreviusSslGrade = getLastDomain.SslGrade
	_, errInsert := services.InsertDomain(domain)
	if errInsert != nil {
		commons.BuilderJSON(w, false, 0, nil)
	}
	commons.BuilderJSON(w, true, http.StatusOK, domain)
}
