package analyzecontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"../commons"
	"../config"
	"../models"
	"github.com/PuerkitoBio/goquery"
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
		domain.SslGrade = getPoorSslGrade(servers)
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
	remote := fmt.Sprintf("http://www.%s", host)
	title := getTokenPage(remote, "title")
	logo := getTokenPage(remote, "link")
	return title, logo, nil
}

func getTokenPage(remote string, token string) string {
	doc, errDom := goquery.NewDocument(remote)
	if errDom != nil {
		log.Fatal(errDom)
	}
	var typeToken string
	var values []string
	doc.Find(token).Each(func(index int, item *goquery.Selection) {
		if textContent := item.Text(); len(textContent) > 0 {
			values = append(values, item.Text())
			if len(typeToken) == 0 {
				typeToken = token
			}
		} else {
			resource, _ := item.Attr("href")
			values = append(values, resource)
			if len(typeToken) == 0 {
				typeToken = token
			}
		}
	})
	fmt.Printf("%s\n", typeToken)
	var value string
	for _, text := range values {
		if typeToken == "title" {
			if len(value) == 0 {
				value = text
			}
		} else if typeToken == "link" {
			if len(string(regexp.MustCompile(".png|.jpg").Find([]byte(text)))) > 0 {
				value = text
			}
		}
	}
	fmt.Printf("%s\n", value)
	return value
}

func getPoorSslGrade(servers []models.Server) string {
	var grades []string
	for _, grade := range servers {
		grades = append(grades, grade.SslGrade)
	}
	sort.Strings(grades)
	return servers[len(grades)-1].SslGrade
}
