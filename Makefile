SWAGGER_SRC := $(wildcard ./internal/server/*.go)
.PHONY: swagger compose-up compose-down integrational-tests unit-tests

swagger: $(SWAGGER_SRC)
	swag init --parseDependency --parseInternal -g ./cmd/app/main.go -o ./swagger

compose-up:
	docker-compose up --build -d && docker-compose logs -f

compose-down:
	docker-compose down

build:
	go build ./cmd/app/main.go
	./main.exe

integrational-tests:
	go test ./tests/integrational -c -o tests.exe
	./tests.exe

unit-tests:
	go test ./internal/service -cover

res:
	go build ./research/research.go
	./research.exe