build-builder:
	docker build -t builder:go-mod -f Dockerfile .
	docker run --rm --name builder -v $(PWD):/web-app -w /web-app builder:go-mod go build -o main main.go

build-service:
	docker build -t web-app:v3 -f service.Dockerfile .
