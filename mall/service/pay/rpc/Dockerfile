FROM alpine:latest

RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/pay/rpc/pay-rpc .
COPY mall/service/pay/etc/pay.yaml etc/pay.yaml

EXPOSE 9003

CMD ["./pay-rpc"]
