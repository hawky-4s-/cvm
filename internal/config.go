package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	configurationUrl = "github.com/hawky-4s-/cvm-matrix"
	configurationMatrixUrl
)

type Config struct {
	Camundas  []CamundaBPM `json:"camundas"`
	Databases []Database   `json:"databases"`
	Servers   []Server     `json:"servers"`
}

func GetRemoteConfiguration(url, username, password string) (*Config, error) {
	filePath := ".cvm_matrix.json"
	DownloadFile(filePath, url, username, password)
	config, err := GetLocalConfiguration(filePath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func GetLocalConfiguration(file string) (*Config, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}

	config, err := ReadConfigurationFromJSON(content)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func ReadConfigurationFromJSON(jsonConfig []byte) (*Config, error) {
	var cfg Config
	err := json.Unmarshal(jsonConfig, &cfg)
	if err != nil {
		fmt.Printf("unable to decode configuration json: %s", err)
		return nil, err
	}
	return &cfg, nil
}
