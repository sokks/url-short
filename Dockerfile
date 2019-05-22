FROM golang:1.10 as builder

ENV SRC_DIR=/go/src/gophersises/02_url_shortener
ADD . $SRC_DIR
WORKDIR $SRC_DIR
RUN CGO_ENABLED=0 GOOS=linux go build -o urlshort 

FROM alpine:3.7

WORKDIR /root
COPY --from=builder /go/src/gophersises/02_url_shortener/urlshort .
# COPY --from=builder /go/src/gitlab.com/sokks/simple-go-app/static ./static
# COPY --from=builder /go/src/gitlab.com/sokks/simple-go-app/data ./data

EXPOSE 12321
ENTRYPOINT [ "./urlshort" ]
