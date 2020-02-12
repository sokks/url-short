FROM nikscorp/go-builder:0.0.1 as builder

ARG SKIP_LINTERS

ENV \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV SRC_DIR=/go/src/github.com/sokks/url-short
ADD . $SRC_DIR
WORKDIR $SRC_DIR
RUN go build -o urlshort
RUN \
    if [ -z "$SKIP_LINTERS" ] ; then \
    golangci-lint run -v ./... ; \
    else echo "skip linters" ; fi

FROM alpine:3.11

WORKDIR /root

COPY --from=builder /go/src/github.com/sokks/url-short/urlshort .
#COPY --from=builder /go/src/github.com/sokks/url-short/wait-for-redis.sh .

EXPOSE 12321
ENTRYPOINT [ "./urlshort" ]
