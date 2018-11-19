FROM golang:1.11 AS builder

RUN mkdir -p /go/src/github.com/alvindcastro/travel-echo-mongo
WORKDIR /go/src/github.com/alvindcastro/travel-echo-mongo

# copies the Gopkg.toml and Gopkg.lock to WORKDIR
COPY ./ ./

# install the dependencies without checking for go code
RUN dep ensure -vendor-only

# Build app
RUN go build -o /app/server .
CMD ["/app/server"]