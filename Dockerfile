FROM golang:1.12.5

WORKDIR $GOPATH/src/go_webapp
COPY . $GOPATH/src/go_webapp

RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go_webapp"]
