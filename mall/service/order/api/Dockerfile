FROM alpine:latest

RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/order/api/order-api .
COPY mall/service/order/api/etc/order.yaml etc/order.yaml

EXPOSE 18002

CMD ["./order-api"]
