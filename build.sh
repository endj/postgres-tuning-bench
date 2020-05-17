rm -f server
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server server.go
docker build -t app .
