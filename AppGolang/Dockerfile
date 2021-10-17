
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .

RUN go mod download && \
    go build -o main .



FROM scratch
COPY --from=builder /build/main /build/index.html ./
EXPOSE 80

ENTRYPOINT ["/main"]


