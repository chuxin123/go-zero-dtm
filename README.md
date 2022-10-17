# go-zero-dtm
go-zero 对接DTM tcc/saga/msg事务demo，模拟订单系统创建订单、库存扣减流程

## 协议使用
所有示例均使用gRPC协议协议编写

## 事务模式
- msg：二阶段消息，适合不需要回滚的全局事务
- saga：适合需要支持回滚的全局事务
- tcc：适合一致性要求较高的全局事务
- xa：适合性能要求不高，没有行锁争抢的全局事务
