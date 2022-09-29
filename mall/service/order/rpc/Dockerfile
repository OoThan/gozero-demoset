FROM alpine:latest

RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/rpc/order-rpc .
COPY mall/service/order/rpc/etc/pay.yaml etc/pay.yaml

EXPOSE 9002

CMD ["./order-rpc"]
