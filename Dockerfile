FROM golang:alpine as build
WORKDIR /go/src/github.com/amsipe/nr-span-example
COPY . .
RUN apk add --no-cache git && \
    go get github.com/newrelic/go-agent && \
    chmod +x wait-for-it.sh && \
    CGO_ENABLED=0 GOOS=linux go build -a -o nr-span-example .
CMD ["./nr-span-example"]
