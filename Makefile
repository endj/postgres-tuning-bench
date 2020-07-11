build:
	rm -f server
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server src/server.go
	docker build -t app .

deploy:
	docker-compose -f deployment/docker-compose.yml up -d 

stop_services:
	-docker stop $$(docker ps -aq) && docker rm $$(docker ps -aq)

ship_it: stop_services build deploy

e2e: ship_it
	 ./test/e2etest.sh sh
