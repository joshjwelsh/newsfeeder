dev:
	go run httpd/main.go

test:
	go test ./...

tcover:
	go test -cover ./...
