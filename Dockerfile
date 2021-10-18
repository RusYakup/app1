FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY ./main.go ./main.go
COPY ./go.mod ./go.mod
COPY ./index.html ./index.html
RUN go build -o main


FROM scratch
WORKDIR /build
COPY --from=builder /build .

EXPOSE 80
ENTRYPOINT ["/build/main"]

