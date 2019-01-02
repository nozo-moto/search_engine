DBNAME:=search_engine
ENV := development
BUILDTARGET := main.go handler.go

all: migrate/up run

dep:
	which dep || go get -v -u github.com/golang/dep/cmd/dep
	dep ensure
	go get github.com/rubenv/sql-migrate/...

run: 
	go run $(BUILDTARGET)
test:
	go test github.com/nozo-moto/search_engine/db

migrate/init:
	mysql -u root -h 127.0.0.1 -P 13306 --protocol tcp -e "create database \`$(DBNAME)\`" -ppassword

migrate/up:
	sql-migrate up -env=$(ENV)

migrate/down:
	sql-migrate down -env=$(ENV)

migrate/status:
	sql-migrate status -env=$(ENV)

docker/start:
	docker-compose up

docker/stop:
	docker-compose stop

make-linux-build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o=bin/main $(BUILDTARGET)

docker-build: make-linux-build
	docker build ./	--tag nozomi0966/search_engine:$(shell git rev-parse --abbrev-ref HEAD | sed -e 's/\//-/g')
