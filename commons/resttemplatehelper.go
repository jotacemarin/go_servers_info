package commons

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"../config"
	"../models"
	"github.com/PuerkitoBio/goquery"
)

// HTTPGet : func
func HTTPGet(url string) ([]byte, error) {
	log.Printf("request to %s", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	return data, nil
}

// GetDomainInfo : func
func GetDomainInfo(host string) (models.DomainR, error) {
	var domainr models.DomainR
	configuration, errConfig := config.LoadConfig()
	er, errHG := HTTPGet(fmt.Sprintf(configuration.SslLabs, host))
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

// GetWhois : func
func GetWhois(endpoints []models.Endpoint) ([]models.Server, error) {
	var servers []models.Server
	for _, endpoint := range endpoints {
		ser1, errSC1 := ShellCall("whois", endpoint.IPAddress, "-A 0 Organization")
		ser2, errSC2 := ShellCall("whois", endpoint.IPAddress, "-A 0 Country")
		if errSC1 != nil {
			return servers, errSC1
		} else if errSC2 != nil {
			return servers, errSC2
		}
		sser1 := strings.TrimSpace(strings.Replace(ser1, "Organization:", "", -1))
		sser2 := strings.TrimSpace(strings.Replace(ser2, "Country:", "", -1))
		servers = append(servers, models.Server{0, endpoint.IPAddress, endpoint.Grade, sser2, sser1})
	}
	return servers, nil
}

// GetPageData : func
func GetPageData(host string) (string, string, error) {
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
	return value
}

// GetPoorSslGrade : func
func GetPoorSslGrade(servers []models.Server) string {
	var grades []string
	for _, grade := range servers {
		grades = append(grades, grade.SslGrade)
	}
	sort.Strings(grades)
	return servers[len(grades)-1].SslGrade
}
