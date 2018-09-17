package artifact

import (
	"github.com/hawky-4s-/clean-http-client"
	"io"

	//"io"
	"net/url"
	"os"
)

func GetBasePartFromUrl(downloadUrl string) (*url.URL, error) {
	parsedUrl, err := url.Parse(downloadUrl)
	if err != nil {
		return nil, err
	}

	return parsedUrl, nil
}

func NewHttpClientFromUrl(downloadUrl, username, password string) (*http.HttpClient, error) {
	// parse url
	parsedUrl, err := GetBasePartFromUrl(downloadUrl)
	if err != nil {
		return nil, err
	}

	// Get the data
	config := http.NewHttpConfig(parsedUrl.Scheme, username, password, "")
	httpClient := http.NewHttpClient(config)

	return httpClient, nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
// https://golangcode.com/download-a-file-from-a-url/
func DownloadFile(filepath, downloadUrl, username, password string) error {
	httpClient, err := NewHttpClientFromUrl(downloadUrl, username, password)
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	basePartFromUrl, _ := GetBasePartFromUrl(downloadUrl)
	resp, err := httpClient.Path(basePartFromUrl.Path).Get().Exec()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
