package analyzecontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../commons"
	"../config"
	"../models"
)

// Index : funct
func Index(w http.ResponseWriter, r *http.Request) {
	er, errDI := getDomainInfo(r.URL.Query().Get("host"))
	information, errW := getWhois(er.Endpoints)
	if errDI != nil || errW != nil {
		commons.BuilderJSON(w, false, 0, nil)
	}
	log.Printf("%s", information)
	commons.BuilderJSON(w, true, http.StatusOK, er)
}

func getDomainInfo(host string) (models.DomainR, error) {
	var domainr models.DomainR
	configuration, errConfig := config.LoadConfig()
	er, errHG := commons.HTTPGet(fmt.Sprintf(configuration.SslLabs, host))
	if errConfig != nil || errHG != nil {
		return domainr, errConfig
	}
	json.Unmarshal([]byte(er), &domainr)
	return domainr, nil
}

func getWhois(endpoints []models.Endpoint) ([]models.Domain, error) {
	var domain []models.Domain
	for _, endpoint := range endpoints {
		ser, errSC := commons.ShellCall("whois", endpoint.IPAddress)
		if errSC != nil {
			return domain, errSC
		}
		log.Printf("-----------\n%s", endpoint.IPAddress)
		log.Printf("\n%s", ser)
	}
	return domain, nil
}
