package internal

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type CamundaBPM struct {
	Version             string `json:"version"`
	MinimalSupportedJDK string `json:"minimal_supported_jdk"`
	DownloadUrl         string `json:"download_url"`
}

func (c CamundaBPM) getDownloadUrl() string {
	return ""
}

func (c *CamundaBPM) isEE() bool {
	return IsCamundaEEVersion(c.Version)
}

func IsCamundaEEVersion(version string) bool {
	return !strings.HasSuffix(strings.ToLower(version), "-ee")
}

// Database is responsible for the database settings and requirements
type Database struct {
	Vendor                      string
	Version                     string
	SupportsBatchProcessing     bool
	JdbcDriverUrl               string
	SupportedSinceCamBPMVersion string
	SupportedUntilCamBPMVersion string
}

func (d *Database) GetConnectionString(dbHost, dbPort string) string {
	return ""
}

type JDK struct {
	Vendor  string `json:"vendor"`
	Version string `json:"version"`
}

func GetMinimalJDK(camunda CamundaBPM, database Database, server Server) {
	//camundaJdk := camunda.MinimalSupportedJDK
	//serverJdk := server.MinimalSupportedJDK
}

type Server struct {
	Name                ServerType `json:"name"`
	Version             string     `json:"version"`
	DownloadUrl         string     `json:"download_url"`
	MinimalSupportedJDK string     `json:"minimal_supported_jdk"`
}

func (s *Server) configureDatabase(database *Database) {

}

type CvmContext struct {
	ProjectName string
	CamundaBPM  CamundaBPM
	Server      Server
	Database    Database
	JDK         JDK
	Backend     Backend
	Properties  map[string]interface{}
	Log         *log.Logger
}

func NewCvmContext(logger *log.Logger) *CvmContext {
	ctx := new(CvmContext)
	ctx.Log = logger

	return ctx
}
