FROM golang:1.17-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokend/notifications/email-mailjet-svc
COPY vendor .
COPY . .

RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux go build  -o /usr/local/bin/email-mailjet-svc gitlab.com/tokend/notifications/email-mailjet-svc


FROM alpine:3.14.6

COPY --from=buildbase /usr/local/bin/email-mailjet-svc /usr/local/bin/email-mailjet-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["email-mailjet-svc"]

