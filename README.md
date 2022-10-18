# go-zero-dtm
go-zero 对接DTM tcc/saga/msg事务示例，模拟订单系统创建订单、库存扣减流程

## 1.协议使用
所有示例均使用gRPC协议编写

## 2.安装及运行
### 2.1 启动DTM
- dtm.yml配置
```
MicroService:
  Driver: 'dtm-driver-gozero' # 配置dtm使用go-zero的微服务协议
  Target: 'etcd://localhost:2379/dtmservice' # 把dtm注册到etcd的这个地址
  EndPoint: 'localhost:36790' # dtm的本地地址
```
- 启动DTM服务
```
dtm -c dtm.yml
```

### 2.2 商品服务gRPC
```
- goctl rpc protoc product.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
- go run product.go -f etc/product.yaml
```

### 2.3 订单服务gRPC
```
- goctl rpc protoc order.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
- go run order.go -f etc/order.yaml
```

### 2.4 订单服务HTTP
```
- goctl api go -api order.api -dir .
- go run order.go -f etc/order.yaml
```

## 3.调试
```
- grpcurl -plaintext 127.0.0.1:8081 product.Product list
- grpcurl -plaintext -d '{"ProductId":1,"Nums":3}' 127.0.0.1:8081 product.Product.checkStock
```

## 4.事务模式
- msg：二阶段消息，适合不需要回滚的全局事务
- saga：适合需要支持回滚的全局事务
- tcc：适合一致性要求较高的全局事务
- xa：适合性能要求不高，没有行锁争抢的全局事务
