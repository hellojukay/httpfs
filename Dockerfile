FROM golang:1.16 AS builder 
RUN mkdir /build
ADD . /build/
WORKDIR /build/
ENV CGO_ENABLED=0
RUN go build -o httpfs

FROM alpine:3.12.1
RUN apk --no-cache add ca-certificates
RUN mkdir /opt/httpfs
WORKDIR /opt/httpfs
COPY --from=builder /build/httpfs /bin/
CMD ["httpfs","-path","/opt/httpfs"]
