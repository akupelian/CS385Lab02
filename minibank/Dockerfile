From alpine as builder
RUN apk update && apk add git && apk add go gcc libc-dev
RUN mkdir ~/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH
COPY * /go/src/
WORKDIR $GOPATH/src
RUN go get -t minibank
RUN go install minibank
WORKDIR ../bin
CMD ["./minibank"]
