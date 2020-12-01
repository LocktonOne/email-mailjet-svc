FROM golang:1.12

WORKDIR /go/src/gitlab.com/tokend/notifications/email-mailjet-svc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/email-mailjet-svc gitlab.com/tokend/notifications/email-mailjet-svc

###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/email-mailjet-svc /usr/local/bin/email-mailjet-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["email-mailjet-svc"]
