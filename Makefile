all: http_flood/http_flood http_flood_client/http_flood_client

http_flood/http_flood: http_flood/http_flood.go
	cd http_flood && go build

http_flood_client/http_flood_client: http_flood_client/client.go
	cd http_flood_client/ && go build

run: http_flood/http_flood
	./http_flood/http_flood
