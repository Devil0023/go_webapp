FROM scratch

WORKDIR $GOPATH/src/go_webapp
COPY . $GOPATH/src/go_webapp

EXPOSE 8000
ENTRYPOINT ["./go_webapp"]


# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go_webapp .