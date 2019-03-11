FROM alpine
RUN apk add --update --no-cache ca-certificates
COPY gateway /dcode-gateway
EXPOSE 4000
ENTRYPOINT ["/dcode-gateway"]