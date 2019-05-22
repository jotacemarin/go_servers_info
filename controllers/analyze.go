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
	information, err := getWhois(r.URL.Query().Get("host"))
	if err != nil {
		commons.BuilderJSON(w, false, 0, nil)
	}
	commons.BuilderJSON(w, true, http.StatusOK, information)
}

func getDomainInfo(host string) (models.DomainR, error) {
	configuration, errConfig := config.LoadConfig()
	if errConfig != nil {
		log.Panicf("error: %s", errConfig.Error())
	}
	var domainr models.DomainR
	er, err := commons.HTTPGet(fmt.Sprintf(configuration.SslLabs, host))
	if err != nil {
		return domainr, err
	}
	json.Unmarshal([]byte(er), &domainr)
	return domainr, nil
}

func getWhois(host string) (models.Domain, error) {
	var domain models.Domain
	er, errHTTPGet := getDomainInfo(host)
	if errHTTPGet != nil {
		log.Panicf("error: %s", errHTTPGet.Error())
		return domain, errHTTPGet
	}
	for _, endpoint := range er.Endpoints {
		log.Printf("%s", endpoint.IPAddress)
		ser, errShellCall := commons.ShellCall(fmt.Sprintf("whois %s", endpoint.IPAddress))
		if errShellCall != nil {
			log.Panicf("error: %s", errShellCall.Error())
			return domain, errShellCall
		}
		log.Printf("%s", ser)
		return domain, nil
	}
	return domain, nil
}
