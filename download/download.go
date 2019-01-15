package download

import (
	"encoding/json"
	"github.com/logrhythm/salt-auto-update/responses"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

func CheckForBadDownload(downloadedFileName string){
	artifactoryErrors := responses.ArtifactoryErrors{}
	downloadedFile, err := ioutil.ReadFile(downloadedFileName)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(downloadedFile, &artifactoryErrors)
	if err == nil {
		if len(artifactoryErrors.Errors) != 0 {
			log.Fatalf("Errors found in downloaded file: %v", artifactoryErrors)
		}
	}
}

func DownloadFile(directory string, fileName string, url string) error {

	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(path.Join(directory, fileName))
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	err = os.Chmod(path.Join(directory, fileName), 0755)
	if err != nil {
		return err
	}

	return nil
}
