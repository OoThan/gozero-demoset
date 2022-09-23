#!/bin/sh

go run user/rpc/user.go -f user/rpc/etc/user.yaml
go run user/api/user.go -f user/api/etc/user.yaml

#go run product/rpc/product.go -f product/rpc/etc/product.yaml
#go run product/api/product.go -f product/api/etc/product.yaml

#go run order/rpc/order.go -f order/rpc/etc/order.yaml
#go run order/api/order.go -f order/api/etc/order.yaml