# === Build stage === #
FROM golang:1.24-alpine3.21 AS builder
COPY go.mod go.sum /go/src/
WORKDIR /go/src/
RUN go mod download && apk add make
COPY . /go/src/
RUN make build name=user_auth ver=1


# === Final stage === #
FROM golang:1.24-alpine3.21

EXPOSE 9999

COPY --from=builder /go/src/cmd/user_auth/v1/user_auth /go/bin/
WORKDIR /go/
RUN adduser -S user
USER user

CMD [ "bin/user_auth" ]
