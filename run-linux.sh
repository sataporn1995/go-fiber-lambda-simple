rm -rf build
export GO111MODULE=on
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o build/main main.go
./build/main