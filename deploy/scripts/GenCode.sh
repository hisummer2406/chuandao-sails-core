# 生成API代码
goctl api go -api delivery-callback.api -dir ../ --style=gozero --home $GOPATH/src/chuandao-sails-core/deploy/goctl/v1.8.5

# 生成RPC代码
# -m 分组
goctl rpc protoc platform.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ -m