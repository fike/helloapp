FROM golang:latest

LABEL MAINTAINER "Fernando Ike <fike@midstorm.org>"

COPY helloapp.go .

RUN CGO_ENABLE=0 go build -o helloapp .

FROM alpine:latest

COPY --from=0 $HOME/go/helloapp /usr/bin/

ENTRYPOINT  /usr/bin/helloapp
