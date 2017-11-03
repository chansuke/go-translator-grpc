build:
	GOOS=linux go build -o app
	docker build -t chansuke/go-translator-grpc .
	rm -f app
