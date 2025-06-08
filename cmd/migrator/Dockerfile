# === Build stage === #
FROM golang:1.24-alpine3.21 AS builder
COPY go.mod go.sum /go/src/
WORKDIR /go/src/
RUN go mod download && apk add make
COPY . /go/src/
RUN make build_migrator

# === Final stage === #
FROM golang:1.24-alpine3.21

COPY --from=builder /go/src/migrator /go/bin/
WORKDIR /go/
RUN adduser -S user
USER user

CMD [ "bin/migrator" ]
