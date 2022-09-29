FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache tzdata && apk add --no-cache ca-certificates
RUN apk add --no-cache libc6-compat
ENV TZ Asia/Shanghai

WORKDIR /app

COPY mall/build/user/api/user-api .
COPY mall/build/user/api/etc/user.yaml etc/user.yaml

EXPOSE 8000

CMD ["./user-api"]
