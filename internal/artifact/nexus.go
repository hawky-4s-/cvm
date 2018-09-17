package artifact

import (
	"fmt"
	"strings"
)

type GAVCoordinate struct {
	Repository string
	GroupId    string
	ArtifactId string
	Version    string
}

func ConstructDownloadUrlFromGAVCoordinates(nexusUrl string, gav *GAVCoordinate) string {
	return fmt.Sprintf("%s?r=%s&g=%s&a=%s&v=%s&p=tar.gz", nexusUrl, gav.Repository, gav.GroupId, gav.ArtifactId, gav.Version)
}

func CalculateGAVCoordinates(server Server, camundabpm CamundaBPM) *GAVCoordinate {
	//distro := server.Name.String()
	repository := "camunda-bpm"
	groupIdTemplate := "org.camunda.bpm.%s"
	artifactIdTemplate := "camunda-bpm"
	//camunda-bpm-ee-{distro}
	// distro one of Tomcat, Wildfly 8 / Rest, JBoss,

	artifactId := artifactIdTemplate

	// if ee version
	if IsCamundaEEVersion(camundabpm.Version) {
		repository += "-ee"
		artifactId = "camunda-bpm-ee"
	}
	// if snapshot wanted
	if strings.HasSuffix(camundabpm.Version, "-SNAPSHOT") {
		repository += "-snapshots"
	}

	return &GAVCoordinate{
		Repository: repository,
		GroupId:    groupIdTemplate,
		ArtifactId: artifactId,
		Version:    camundabpm.Version,
	}
}

func DownloadCamundaFromNexusToFile(file string, nexusUrl string, username string, password string, server Server, camundabpm CamundaBPM) error {
	gavCoordinate := CalculateGAVCoordinates(server, camundabpm)
	downloadUrlFromGAVCoordinates := ConstructDownloadUrlFromGAVCoordinates(nexusUrl, gavCoordinate)
	return DownloadFile(file, downloadUrlFromGAVCoordinates, username, password)
}

//# Determine nexus URL parameters
//if [ ${EE} = "true" ]; then
//echo "Downloading Camunda ${VERSION} Enterprise Edition for ${DISTRO}"
//REPO="camunda-bpm-ee"
//ARTIFACT="camunda-bpm-ee-${DISTRO}"
//ARTIFACT_VERSION="${VERSION}-ee"
//else
//echo "Downloading Camunda ${VERSION} Community Edition for ${DISTRO}"
//REPO="camunda-bpm"
//ARTIFACT="camunda-bpm-${DISTRO}"
//ARTIFACT_VERSION="${VERSION}"
//fi
//
//# Determine if SNAPSHOT repo and version should be used
//if [ ${SNAPSHOT} = "true" ]; then
//REPO="${REPO}-snapshots"
//ARTIFACT_VERSION="${VERSION}-SNAPSHOT"
//fi
//
//# Determine artifact group, all wildfly version have the same group
//case ${DISTRO} in
//wildfly*) GROUP="wildfly" ;;
//*) GROUP="${DISTRO}" ;;
//esac
//ARTIFACT_GROUP="org.camunda.bpm.${GROUP}"

func GetAvailableCamundaBPMVersions(username string, password string) {
	if username != "" && password != "" {

	}
}

// https://github.com/camunda-ci/camunda-ci-maintenance/blob/master/nexus-cli/src/main/groovy/nexus/NexusCli.groovy
func constructGAVQuery() {}
