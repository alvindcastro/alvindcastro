FROM golang 
WORKDIR /Users/alvin/workspaces/go/src/github.com/alvindcastro/travel-rest-api
COPY . .
RUN set -x && \ 
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o travel-rest-api .

FROM scratch
WORKDIR /root/
COPY --from=0 /Users/alvin/workspaces/go/src/github.com/alvindcastro/travel-rest-api .
EXPOSE 3000
CMD ["./travel-rest-api"]  
