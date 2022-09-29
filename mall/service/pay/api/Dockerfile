FROM alpine:latest

RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/pay/api/pay-api .
COPY mall/service/pay/api/etc/pay.yaml etc/pay.yaml

EXPOSE 18003

CMD ["./pay-api"]
