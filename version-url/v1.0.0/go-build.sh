CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o version-tool-win-amd64.exe main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o version-tool-win-arm64.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o version-tool-mac-amd64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o version-tool-mac-arm64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o version-tool-linux-arm64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o version-tool-linux-amd64 main.go