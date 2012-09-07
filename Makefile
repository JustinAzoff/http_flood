all: http_flood

http_flood: http_flood.go
	go build http_flood.go

run: http_flood
	./http_flood
