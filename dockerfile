FROM golang:1.16.3 as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY .  ./
# ADD .  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go env -w GOPROXY=direct GOFLAGS="-insecure"
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'


FROM scratch as runner

COPY --from=builder /go/bin/main /app/main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/main"]