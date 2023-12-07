rm -rf build
export GO111MODULE=on
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o build/main main.go
zip -jrm build/main.zip build/main