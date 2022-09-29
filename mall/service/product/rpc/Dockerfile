FROM alpine:latest

RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/product/rpc/product-rpc .
COPY mall/service/product/rpc/etc/product.yaml etc/product.yaml

EXPOSE 9001

CMD ["./product-rpc"]
