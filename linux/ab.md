# ab 测试

## 安装

```shell
sudo apt install apache2-utils
```

## 使用

```shell
ab -c 10 -n 1000 http://localhost:8080/xxx
```

## 参数说明

* -n 执行次数
* -c 并发数

## 指标

* Requests per second: 吞吐量，某个并发用户数下单位时间内处理的请求数
* Concurrency Level: 并发数量 `-c`
* Complete requests: 总请求书 `-n`
* Failed requests: 错误数量
* Total transferred: 所有请求的响应数据长度总和
* HTML transferred: 所有请求的响应数据中正文数据的总和
* Time per request: 用户平均请求等待时间
* Time per request: across all concurrent requests 服务器平均等待时长
* Time taken for tests: 总时长
* Transfer rate: 单位时间内从服务器获取的数据长度
