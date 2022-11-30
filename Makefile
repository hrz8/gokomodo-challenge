run:
	go build -o bin/server cmd/server/main.go
	./bin/server

lint:
	golangci-lint run

test:
	go test -v ./...

docker-build:
	docker build --tag gokomodo-challenge:1.0 .

docker-start:
	docker create --name gokomodo-app -p 3000:3000 gokomodo-challenge:1.0
	docker start gokomodo-app
	docker logs gokomodo-app -f

docker-compose:
	docker-compose up --build -d
	docker-compose logs -f
