FROM golang
COPY . /go/src/otl
WORKDIR /go/src/otl
RUN go get .
ENTRYPOINT go run server.go
EXPOSE 9000