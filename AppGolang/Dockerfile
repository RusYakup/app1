

FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN apt-get update
RUN apt-get install Sudo


FROM alpine
RUN apt-get update && apt-get install -y \
curl
CMD /bin/bash


FROM scratch
COPY --from=builder /build/main /build/index.html ./

EXPOSE 80
ENTRYPOINT ["/main"]

