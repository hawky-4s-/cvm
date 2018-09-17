package internal

import "testing"

func TestReadConfigurationFromJSON(t *testing.T) {
	configJson := []byte(`
{
    "camundas" : [
        {}
    ],
    "databases" : [
        {}
    ],
    "servers" : [
        {
            "name": "tomcat",
            "version": "9.0.5"
        }
    ]
}
`)
	config, err := ReadConfigurationFromJSON(configJson)
	if err != nil {
		t.Fatal(err)
	}
	if config == nil {
		t.Fatal("config must not be nil")
	}
}

func TestReadLocalConfiguration(t *testing.T) {

}

func TestGetRemoteConfiguration(t *testing.T) {

}
