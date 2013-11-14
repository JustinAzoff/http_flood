all: static http_flood http_flood_client/http_flood_client

static: web/index.html.go web/Flashflood.swf.go

web/Flashflood.swf: web/Flashflood.hx
	cd web && make

web/index.html.go: web/index.html
	rm -f web/index.html.go
	go-bindata -f Index_html -p web -i web/index.html

web/Flashflood.swf.go: web/Flashflood.swf
	rm -f web/Flashflood.swf.go
	go-bindata -f Flashflood_swf -p web -i web/Flashflood.swf

http_flood: http_flood.go common/common.go consts/consts.go web/index.html.go web/Flashflood.swf.go
	go build http_flood.go

http_flood_client/http_flood_client: http_flood_client/client.go common/common.go consts/consts.go
	cd http_flood_client/ && go build

run: http_flood
	./http_flood
