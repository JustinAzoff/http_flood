all: http_flood client

http_flood: http_flood.go
	go build http_flood.go

client: client.go
	go build client.go

run: http_flood
	./http_flood
