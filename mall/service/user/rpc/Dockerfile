FROM alpine:latest

RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/user/rpc/user-rpc .
COPY mall/service/user/rpc/etc/user.yaml etc/user.yaml

EXPOSE 9000

CMD ["./user-rpc"]
