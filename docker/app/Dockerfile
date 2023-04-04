FROM golang:1.16-alpine

WORKDIR /go/src
COPY ./src .

RUN  apk update && \ 
        apk add git && \
        go get github.com/cosmtrek/air@v1.29.0

CMD ["air", "-c", ".air.toml"]
