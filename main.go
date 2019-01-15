package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logrhythm/salt-auto-update/config"
	"github.com/logrhythm/salt-auto-update/download"
	"github.com/logrhythm/salt-auto-update/responses"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path"
	"regexp"
)

const WebBundleInstallerName = "LRWebServices.exe"
const BaseArtifactoryUrl = "http://usbo1part01:8081/artifactory/STABLE/web-services-bundle/8.0.0/win/LRWebServices_64_"

func main() {
	configuration := config.Config{}

	configurationFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(configurationFile, &configuration)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	resp, err := http.Get("http://usbo1part01.schq.secious.com:8081/artifactory/api/storage/STABLE/web-services-bundle/8.0.0/win/?lastModified")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	artifactoryLatestVersion := responses.ArtifactoryLatestVersion{}
	err = json.Unmarshal(respBytes, &artifactoryLatestVersion)
	if err != nil {
		log.Fatal(err)
	}

	reg, err := regexp.Compile(`[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*`)
	if err != nil {
		log.Fatal(err)
	}
	webBundleVersion := reg.FindString(artifactoryLatestVersion.Uri)

	fmt.Printf("Downloading version %v to %v.\n", webBundleVersion, configuration.SaltInstallerDirectory)

	finalArtifactoryUrl := fmt.Sprintf("%v%v.exe", BaseArtifactoryUrl, webBundleVersion)

	err = download.DownloadFile(configuration.SaltInstallerDirectory, WebBundleInstallerName, finalArtifactoryUrl)
	if err != nil {
		log.Fatal(err)
	}

	download.CheckForBadDownload(path.Join(configuration.SaltInstallerDirectory,WebBundleInstallerName))

	RunSaltInstall()
}

func RunSaltInstall() {
	cmd := exec.Command("/bin/sh", "salt-script.sh")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
