package analyzecontroller

import (
	"log"
	"net/http"

	"../commons"
	"../models"
)

// Index : funct
func Index(w http.ResponseWriter, r *http.Request) {
	er, errDI := commons.GetDomainInfo(r.URL.Query().Get("host"))
	if errDI != nil {
		commons.BuilderJSON(w, false, 0, nil)
	}
	var domain models.Domain
	if title, logo, errP := commons.GetPageData(r.URL.Query().Get("host")); len(title) > 0 || len(logo) > 0 {
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
	commons.BuilderJSON(w, true, http.StatusOK, domain)
}
