FROM golang:1.18-alpine3.15

EXPOSE 9000

RUN apk update \
  && apk add --no-cache \
    mysql-client \
    build-base

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY ./hex_entrypoint.sh /usr/local/bin/hex_entrypoint.sh
RUN /bin/chmod +x /usr/local/bin/hex_entrypoint.sh

RUN go build cmd/main.go
RUN mv main /usr/local/bin

ENTRYPOINT [ "hex_entrypoint.sh" ]

CMD [ "main" ]