http_flood
==========

http server that floods the client with random data. It has endpoints for uploading and downloading data.

It logs information like

    2013/05/22 20:03:14 Listening on :7070
    2013/05/22 20:03:17 flood starting addr=127.0.0.1:50629 megabytes=10240
    2013/05/22 20:03:21 flood finished addr=127.0.0.1:50629 duration=4.107445s megabytes=10240.0 speed=2493.0MB/s
    2013/05/22 20:05:18 upload starting addr=127.0.0.1:50632
    2013/05/22 20:05:23 upload finished addr=127.0.0.1:50632 duration=4.999599s megabytes=4244.6 speed=849.0MB/s

Building
========

    go get github.com/jteeuwen/go-bindata
    make

Usage
=====

Start http_flood:

    $ ./http_flood
    2013/05/22 20:04:09 Listening on :7070

Hit it with http_flood_client

    $ ./http_flood_client/http_flood_client 
    2013/05/22 20:04:43 download finished duration=10.010801s megabytes=22094.3 speed=2207.1MB/s
    2013/05/22 20:04:53 upload finished addr=127.0.0.1:50630 duration=9.999803s megabytes=19622.1 speed=1962.3MB/s

Run a full duplex test:

    $ ./http_flood_client/http_flood_client --full --seconds 5
    2013/05/22 20:05:23 upload finished addr=127.0.0.1:50632 duration=4.999599s megabytes=4244.6 speed=849.0MB/s
    2013/05/22 20:05:23 download finished duration=5.01283s megabytes=3859.3 speed=769.9MB/s

Test using wget:

    $ wget -O/dev/null localhost:7070/flood?m=10240
    --2013-05-22 20:06:08--  http://localhost:7070/flood?m=10240
    Resolving localhost (localhost)... 127.0.0.1
    Connecting to localhost (localhost)|127.0.0.1|:7070... connected.
    HTTP request sent, awaiting response... 200 OK
    Length: 10737418240 (10G) [application/octet-stream]
    Saving to: ‘/dev/null’

    100%[===========================>] 10,737,418,240 2.49GB/s   in 4.0s   

    2013-05-22 20:06:12 (2.49 GB/s) - ‘/dev/null’ saved [10737418240/10737418240]


API
===

* GET /flood?m=1024  -- Download 1024 megabytes
* GET /flood?s=10    -- Download for 10 seconds
* POST /upload       -- Upload arbitrary data
