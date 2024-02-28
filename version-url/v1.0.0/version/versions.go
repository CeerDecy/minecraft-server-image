package version

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const VersionManifestRequestUrl = "https://launchermeta.mojang.com/mc/game/version_manifest.json"

type Latest struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

type Versions struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Time        string `json:"time"`
	ReleaseTime string `json:"releaseTime"`
}

type VersionManifest struct {
	Latest   Latest     `json:"latest"`
	Versions []Versions `json:"versions"`
}

func GetVersionManifest() (*VersionManifest, error) {
	var manifest VersionManifest
	request, err := http.NewRequest(http.MethodGet, VersionManifestRequestUrl, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &manifest, json.Unmarshal(body, &manifest)
}

func GetVersionUrl(version string) (string, error) {
	manifest, err := GetVersionManifest()
	if err != nil {
		return "", err
	}
	for _, v := range manifest.Versions {
		if v.Id == version {
			return v.Url, nil
		}
	}
	return "", fmt.Errorf("can't find version: [%v]", version)
}
