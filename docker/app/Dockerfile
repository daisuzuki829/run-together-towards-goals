#Step 1 Builld
FROM golang:alpine as builder

WORKDIR /go/src/app/
ADD . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o app .

#Step2 Move build file
FROM alpine:latest

RUN apk add --no-cache git vim less curl make gcc mariadb-client mysql-dev alpine-sdk

COPY --from=builder /go/src/app/ /app

ENV PORT=${PORT}

WORKDIR /app

ENTRYPOINT ["./app"]
