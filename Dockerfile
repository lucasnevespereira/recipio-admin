FROM golang:1.20-alpine3.17

RUN apk add -v build-base
RUN apk add -v ca-certificates
RUN apk add --no-cache \
    unzip \
    openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o recipio-admin

EXPOSE 8080

CMD ["/app/recipio-admin", "serve", "--http=0.0.0.0:8080"]