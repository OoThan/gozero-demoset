FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache tzdata && apk add --no-cache ca-certificates
RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/product/rpc/product-rpc .
COPY mall/build/product/rpc/etc/product.yaml etc/product.yaml

RUN ls -al
RUN pwd

EXPOSE 9000

CMD ["./product-rpc"]
