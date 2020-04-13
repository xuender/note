# kafka 消息队列

## 工作原理

* 发送消息
    1. 发送消息到kafka
    1. kafka追加消息到消息日志末尾
* 接收消息
    1. 访问zookeeper获取kafka服务器地址
    1. 访问zookeeper获取消息队列位置
    1. 根据位置获取kafka保存的对应消息

## 注意事项

kafka配置中要说明ADVERTISED_HOST，这个变量告诉zookeeper访问呢kafka时使用的地址或域名，默认是localhost
如果不配置这个变量，就不能在其他机器访问kafka

## docker 配置

docker-compose.yml

```yml
version: '3'
services:
    zookeeper:          # 服务协调
        container_name: zookeeper # 容器名称
        image: wurstmeister/zookeeper # 镜像
        restart: always # 重启继续运行
        network_mode: host # 网路模式
        ports:
            - 2181:2181
        environment:
            ZOOKEEPER_CLIENT_PORT: 2181
            ZOOKEEPER_TICK_TIME: 2000
    kafka:          # 消息队列
        container_name: kafka
        image: wurstmeister/kafka
        restart: always
        network_mode: host
        ports:
            - 9092:9092
        volumes:
            - /tmp/kafka:/kafka # kafka日志文件保存目录
        environment:
            KAFKA_ZOOKEEPER_CONNECT: 192.168.1.12:2181
            KAFKA_LOG_DIRS: /kafka/logs
            KAFKA_LISTENERS: PLAINTEXT://:9092
            KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://192.168.1.12:9093
            KAFKA_ADVERTISED_HOST_NAME: 192.168.1.12 # 访问zookeeper时返回的kafka地址
            KAFKA_LOG_RETENTION_HOURS: 48 # 日志保存48小时
            KAFKA_LOG_RETENTION_BYTES: 1073741824 # 日志保存1G
        depends_on:
            - zookeeper # 依赖
```

## 常用命令

使用控制台创建消息

```bash
bin/kafka-console-producer.sh --broker-list 192.168.1.12:9092 --topic test
```

读取控制台输入的消息

```bash
bin/kafka-console-consumer.sh --zookeeper 192.168.1.12:2181 --topic test --from-beginning
```
