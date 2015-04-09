# http_flood
#
# VERSION               0.1

FROM google/golang
MAINTAINER Justin Azoff <justin.azoff@gmail.com>

WORKDIR /gopath/src/github.com/JustinAzoff/http_flood/
ADD . /gopath/src/github.com/JustinAzoff/http_flood/

RUN go get
RUN go build
RUN find /gopath

EXPOSE 7070
ENTRYPOINT ["/gopath/bin/http_flood", "server"]
