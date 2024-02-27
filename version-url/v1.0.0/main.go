package main

import (
	"flag"
	"fmt"

	pkg "version-url/package"
	"version-url/version"
)

func GetDownloadUrl(ver string) (string, error) {
	url, err := version.GetVersionUrl(ver)
	if err != nil {
		return "", err
	}
	versionPackage, err := pkg.GetVersionPackage(url)
	if err != nil {
		return "", err
	}
	return versionPackage.Downloads.Server.Url, nil
}

func main() {
	var v = flag.String("version", "", "The version of minecraft")
	flag.Parse()
	url, err := GetDownloadUrl(*v)
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	fmt.Println("ok", url)
}
