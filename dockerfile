# used to ssl certificate
FROM debian:latest AS cert
RUN apt update && apt install ssl-cert -y

FROM golang:1.15 AS build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o edgex-restapp ./cmd/edgex-restapp-server/

FROM scratch
COPY --from=cert /etc/ssl/certs/ssl-cert-snakeoil.pem tls/certificate.crt
COPY --from=cert /etc/ssl/private/ssl-cert-snakeoil.key tls/key.key
COPY --from=build /app/edgex-restapp .

VOLUME [ "/persist" ]

EXPOSE 8080
EXPOSE 8443

ENV TZ=Europe/Berlin
CMD ["/edgex-restapp"]