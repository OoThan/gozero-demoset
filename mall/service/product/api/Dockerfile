FROM alpine:latest

RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/product/api/product-api .
COPY mall/service/product/api/etc/product.yaml etc/product.yaml

EXPOSE 18001

CMD ["./product-api"]
