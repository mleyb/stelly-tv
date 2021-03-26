go get -v

$env:GOOS="windows";$env:GOARCH="amd64"; go build -o .\bin\twitch-authenticate.exe

$env:GOOS="linux"; go build -o .\bin\twitch-authenticate