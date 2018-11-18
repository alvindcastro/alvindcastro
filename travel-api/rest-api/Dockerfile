FROM golang:1.11 as builder
WORKDIR $GOPATH/src/github.com/alvindcastro/travel-api
COPY . .
RUN set -x && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o travel-api .

FROM scratch
WORKDIR /root/
COPY --from=0 $GOPATH/src/github.com/alvindcastro/travel-api .
EXPOSE 3000
CMD ["./travel-api"]
