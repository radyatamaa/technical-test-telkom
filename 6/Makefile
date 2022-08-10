BINARY=engine
commit:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	git commit

test: 
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

engine:
	go build -o ${BINARY} main.go

unittest:
	go test -short  ./...

docker:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	docker build --pull --rm -f "Dockerfile" -t technical-test-telkom "." --build-arg BRANCH=develop

run:
	docker-compose up -d

stop:
	docker-compose down
