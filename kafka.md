- Kfaka 
* 吞吐量，性能
* 伸缩性
* 容错性
* 大数据结合

- 概念：

- topic
- partition - offset(偏移量)
- record(key, value) - key 决定分区
- 副本机制 保障数据可靠性（leader, follower）

- Broker 消息代理，负责数据读写，一台服务器一个broker 
- 一般流程： broker 注册到 zookeeper 中，zookeepe r 提供给 client 地址

- 分区最小的并行单位
- 消费者可以消费多个分区
- 分区也可以被多个消费者组里的消费消费
- 但是，一个分区不能被同一个消费者里的多个消费者消费

- 顺序： 同一个生成者发送到同一个分区，先发的offset 小
- 分区内顺序有保障，不同分区无法保证
  - 解决方案： 设置相同key

- 消费传递类型： 1.最多一次 2.最少一次 3.精确一次

- 事务消息（隔离级别）

- 序列化
  - csv
  - json - es
  - 序列： avro / protobuf - 大数据
  - schema registry
   - 背景，多模式，顺序要求
  - record： header


[参照链接](https://www.bilibili.com/video/BV1h94y1Q7Xg?spm_id_from=333.788.player.switch&vd_source=d1308681e37e55f59ae344651b974ea4&p=8) 