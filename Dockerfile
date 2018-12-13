FROM golang:alpine3.7 AS build-server
COPY . /usr/local/go/src/gows 

WORKDIR /usr/local/go/src/gows
RUN go get \
    && go build -ldflags "-s -w" 

FROM alpine:3.7
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY --from=build-server /usr/local/go/src/gows/gows .

EXPOSE 9000

CMD ./gows -p=9000 -debug=true
