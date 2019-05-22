FROM golang:1.10 as builder

ENV SRC_DIR=/go/src/github.com/sokks/url-short
ADD . $SRC_DIR
WORKDIR $SRC_DIR
RUN CGO_ENABLED=0 GOOS=linux go build -o urlshort 

FROM alpine:3.7

WORKDIR /root

COPY --from=builder /go/src/github.com/sokks/url-short/urlshort .
COPY --from=builder /go/src/github.com/sokks/url-short/wait-for-redis.sh .

EXPOSE 12321
ENTRYPOINT [ "./urlshort" ]
