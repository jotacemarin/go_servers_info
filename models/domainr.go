package models

// DomainR : struct
type DomainR struct {
	Host            string     `json:"host"`
	Port            int        `json:"port"`
	Protocol        string     `json:"protocol"`
	IsPublic        string     `json:"isPublic"`
	Status          string     `json:"status"`
	StartTime       int        `json:"startTime"`
	TestTime        int        `json:"testTime"`
	EngineVersion   string     `json:"engineVersion"`
	CriteriaVersion string     `json:"criteriaVersion"`
	Endpoints       []Endpoint `json:"endpoints"`
}
