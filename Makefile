DBNAME:=search_engine
ENV := development

dep:
	which dep || go get -v -u github.com/golang/dep/cmd/dep
	dep ensure
	go get github.com/rubenv/sql-migrate/...

run: 
	go run main.go handler.go

test:
	go test github.com/nozo-moto/search_engine/db

migrate/init:
	mysql -u root -h 0.0.0.0 -P 13306 --protocol tcp -e "create database \`$(DBNAME)\`" -p

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

