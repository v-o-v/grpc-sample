export GOCMD=go
export GO111MODULE=on
export GIN_MODE=debug

export PORT=443
export SSL_CRT=./docker/ssl/aiko.pem
export SSL_KEY=./docker/ssl/aiko.key

run:
	go run main.go

debug:
	go run main.go

client:
	go run client/main.go

build:
	docker build -f ./docker/Dockerfile.stg .
