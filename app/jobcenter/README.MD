# 任务中心模块
## 定时任务
定期拉取第三方k线图数据保存到mongo中
## kafka
### 创建主题
进入kafka01容器内执行该指令
```sh

kafka-topics.sh --bootstrap-server kafka01:9092 --topic kline-1m-BTC-USDT --create --partitions 1 --replication-factor 2
kafka-topics.sh --bootstrap-server kafka01:9092 --topic kline-1m-ETH-USDT --create --partitions 1 --replication-factor 2

```