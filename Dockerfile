# syntax=docker/dockerfile:1
FROM golang:1.16
WORKDIR /iban-service
COPY . /iban-service/
RUN CGO_ENABLED=0 GOOS=linux go build -o iban-service app/*.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /iban-service/iban-service .
COPY --from=0 /iban-service/entrypoint.sh .
ENTRYPOINT ["./entrypoint.sh"]
