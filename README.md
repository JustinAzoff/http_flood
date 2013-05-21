http_flood
==========

http server that floods the client with random data

So far, outputs stuff like

    flood starting addr=127.0.0.1:53868
    flood finished addr=127.0.0.1:53868 duration=1.497196s megabytes=1000.0 speed=667.9MB/s

Building
========

    go install github.com/jteeuwen/go-bindata
    make
