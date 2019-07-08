FROM golang:latest

WORKDIR $GOPAH/src/github.com/go_webapp
COPY . $GOPAH/src/github.com/go_webapp
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go_webapp"]
