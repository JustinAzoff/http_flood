# http_flood
#
# VERSION               0.1

FROM      stackbrew/ubuntu:13.04
MAINTAINER Justin Azoff <justin.azoff@gmail.com>

RUN apt-get update
RUN apt-get install -y golang
ADD . http_flood
RUN cd http_flood && go build
EXPOSE 7070
CMD ["http_flood/http_flood"]
