FROM golang:1.15.5 AS builder 
RUN mkdir /build
ADD . /build/
WORKDIR /build/
ENV CGO_ENABLED=0
RUN go build -o httpfs main.go

FROM alpine:3.12.1
RUN apk --no-cache add ca-certificates
RUN mkdir /opt/httpfs
WORKDIR /opt/httpfs
COPY --from=builder /build/httpfs /bin/
CMD ["httpfs","-path","/opt/httpfs"]
