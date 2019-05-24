package models

// Config : struct
type Config struct {
	Port         string `json:"port"`
	SslLabs      string `json:"ssllabs"`
	DbDriver     string `json:"db_driver"`
	DbConnection string `json:"db_connection"`
}
