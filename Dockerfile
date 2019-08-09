FROM scratch

WORKDIR $GOPATH/src/go_webapp
COPY . $GOPATH/src/go_webapp

EXPOSE 8000
ENTRYPOINT ["./go_webapp"]