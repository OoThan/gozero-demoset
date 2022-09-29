FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache tzdata && apk add --no-cache ca-certificates
RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/pay/rpc/pay-rpc .
COPY mall/build/pay/rpc/etc/pay.yaml etc/pay.yaml

RUN ls -al
RUN pwd

EXPOSE 9000

CMD ["./pay-rpc"]
