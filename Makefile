build-builder:
	docker build -t builder:go-mod -f Dockerfile .
	docker run --rm --name builder -v $(PWD):/web-app -w /web-app builder:go-mod go build -o main main.go

build-service:
	docker build -t gomod -f service.Dockerfile .

run-service:
	docker run -d --rm --name gomod -p ${PORT_NUMBER}:${PORT_NUMBER} -e "PORT_NUMBER=${PORT_NUMBER}" gomod
