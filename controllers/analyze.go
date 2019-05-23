package analyzecontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"../commons"
	"../config"
	"../models"
	"golang.org/x/net/html"
)

// Index : funct
func Index(w http.ResponseWriter, r *http.Request) {
	er, errDI := getDomainInfo(r.URL.Query().Get("host"))
	if errDI != nil {
		commons.BuilderJSON(w, false, 0, nil)
	}
	var domain models.Domain
	if title, logo, errP := getPageData(r.URL.Query().Get("host")); len(title) > 0 || len(logo) > 0 {
		if errP != nil {
			log.Print(errP)
			commons.BuilderJSON(w, false, 0, nil)
		}
		domain.Title = title
		domain.Logo = logo
	}
	if endpointsLength := len(er.Endpoints); endpointsLength > 0 {
		servers, errW := getWhois(er.Endpoints)
		if errW != nil {
			commons.BuilderJSON(w, false, 0, nil)
		}
		domain.Servers = servers
		domain.IsDown = false
	} else {
		domain.IsDown = true
	}
	commons.BuilderJSON(w, true, http.StatusOK, domain)
}

func getDomainInfo(host string) (models.DomainR, error) {
	var domainr models.DomainR
	configuration, errConfig := config.LoadConfig()
	er, errHG := commons.HTTPGet(fmt.Sprintf(configuration.SslLabs, host))
	if errConfig != nil || errHG != nil {
		return domainr, errConfig
	}
	json.Unmarshal([]byte(er), &domainr)
	var domain models.Domain
	if endpointsLength := len(domainr.Endpoints); endpointsLength > 0 {
		domain.IsDown = false
	}
	return domainr, nil
}

func getWhois(endpoints []models.Endpoint) ([]models.Server, error) {
	var servers []models.Server
	for _, endpoint := range endpoints {
		ser1, errSC1 := commons.ShellCall("whois", endpoint.IPAddress, "-A 0 Organization")
		ser2, errSC2 := commons.ShellCall("whois", endpoint.IPAddress, "-A 0 Country")
		if errSC1 != nil {
			return servers, errSC1
		} else if errSC2 != nil {
			return servers, errSC2
		}
		sser1 := strings.TrimSpace(strings.Replace(ser1, "Organization:", "", -1))
		sser2 := strings.TrimSpace(strings.Replace(ser2, "Country:", "", -1))
		servers = append(servers, models.Server{endpoint.IPAddress, endpoint.Grade, sser2, sser1})
	}
	return servers, nil
}

func getPageData(host string) (string, string, error) {
	response, errP := commons.HTTPGet(fmt.Sprintf("http://www.%s", host))
	if errP != nil {
		return "", "", errP
	}
	title := getTokenPage(string(response), "title")
	logo := getTokenPage(string(response), "link")
	return title, logo, nil
}

func getTokenPage(response string, token string) string {
	domDoc := html.NewTokenizer(strings.NewReader(response))
	previousToken := domDoc.Token()
	var content []string
loopDoc:
	for {
		tt := domDoc.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDoc // End of the document,  done
		case tt == html.StartTagToken:
			previousToken = domDoc.Token()
		case tt == html.TextToken:
			if previousToken.Data != token {
				continue
			}
			if TxtContent := strings.TrimSpace(html.UnescapeString(string(domDoc.Text()))); len(TxtContent) > 0 {
				fmt.Printf("%s\n", TxtContent)
				content = append(content, TxtContent)
			}
		}
	}
	if length := len(content); length > 0 {
		return content[0]
	} else {
		return ""
	}
}
