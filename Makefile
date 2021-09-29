.PHONY = build test
build:
	go get -u github.com/swaggo/swag/cmd/swag
	swag init -g ./cmd/server/main.go -o docs
	go build -v ./cmd/server

build-docker:
	docker build --no-cache --network host -f ./docker/builder.Dockerfile . --tag patreon

run:
	#sudo chown -R 5050:5050 ./pgadmin
	docker-compose up --build --no-deps

stop:
	docker-compose down

rm-docker:
	docker rm -vf $$(docker ps -a -q) || true

run-coverage:
	go test -covermode=atomic -coverpkg=./... -coverprofile=cover ./...
	cat cover | fgrep -v "mock" | fgrep -v "pb.go" | fgrep -v "testing.go" | fgrep -v "teststore" | fgrep -v "easyjson" | fgrep -v "start.go" > cover2
	go tool cover -func=cover2

test:
	go test -v -race ./...
