FROM alpine:latest

RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/user/api/user-api .
COPY mall/service/user/api/etc/user.yaml etc/user.yaml

EXPOSE 8000

CMD ["./user-api"]
