FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache tzdata && apk add --no-cache ca-certificates
RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/order/rpc/order-rpc .
COPY mall/build/order/rpc/etc/order.yaml etc/order.yaml

EXPOSE 9000

CMD ["./order-rpc"]
