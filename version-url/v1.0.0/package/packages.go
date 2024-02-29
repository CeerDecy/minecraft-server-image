package _package

import (
	"encoding/json"
	"io"
	"net/http"
)

type Package struct {
	Downloads Downloads `json:"downloads"`
}

type Downloads struct {
	Server Server `json:"server"`
}

type Server struct {
	Sha1 string `json:"sha1"`
	Size uint64 `json:"size"`
	Url  string `json:"url"`
}

func GetVersionPackage(url string) (*Package, error) {
	var pkg Package
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp.Body == nil {
			_ = resp.Body.Close()
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &pkg, json.Unmarshal(body, &pkg)
}
