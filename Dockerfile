FROM golang:1.13-alpine as builder
WORKDIR /src/

ENV GO111MODULE=on

COPY . /src

RUN go mod download
RUN go build -o barb cmd/barb/main.go

FROM alpine:latest
WORKDIR /root/

RUN apk add --no-cache tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/barb .

CMD ["./barb"]
