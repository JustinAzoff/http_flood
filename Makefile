all: static http_flood

static: web/index.html.go web/Flashflood.swf.go

web/Flashflood.swf: web/Flashflood.hx
	cd web && make

web/index.html.go: web/index.html
	rm -f web/index.html.go
	go-bindata -func Index_html -pkg web web/index.html

web/Flashflood.swf.go: web/Flashflood.swf
	rm -f web/Flashflood.swf.go
	go-bindata -func Flashflood_swf -pkg web web/Flashflood.swf

http_flood: main.go client.go server.go randomreader/randomreader.go consts/consts.go web/index.html.go web/Flashflood.swf.go
	go build

run: http_flood
	./http_flood server

install: all
	install -d ${DESTDIR}/usr/bin/
	install http_flood http_flood_client/http_flood_client ${DESTDIR}/usr/bin/
