# prometheus

普罗米修斯，监控&报警

## docker 安装

```shell
# 服务器监控
docker pull prom/node-exporter
# 数据抓取
docker pull prom/prometheus
# 显示存储
docker pull grafana/grafana
# 推送接收
docker pull prom/pushgateway
# docker run -d -p 9091:9091 prom/pushgateway
```

Linux机器指标收集

```shell
docker run -d -p 9100:9100 \
  -v "/proc:/host/proc:ro" \
  -v "/sys:/host/sys:ro" \
  -v "/:/rootfs:ro" \
  --net="host" \
  --name=exporter \
  prom/node-exporter
```

prometheus 主服务

```shell
docker run  -d \
  -p 9090:9090 \
  -v /opt/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml  \
  --restart=always \
  --name=prom \
  prom/prometheus
```

图形界面数据存储

```shell
docker run -d \
  -p 3000:3000 \
  -v /opt/grafana-storage:/var/lib/grafana \
  --restart=always \
  --name=grafana \
  grafana/grafana
```

执行

```shell
docker exec -it prom /bin/sh
```

## metric 指标

### Counter (计数器)

单调递增的metrics，比如：服务的请求数、已完成的任务数，错误发生的次数等。

## CounterVec (计数器组)

一组计数器，有不同的纬度,例如：http_requests_total，"code", "method"

### Gauge (仪表盘)

数据可以任意变化的metrics，一般用于像温度或者内存使用率这种数据(任意上下变化)，当然也用于表示可以上下变化的"总数"，比如：当前运行的goroutines总数。

### GaugeVec (仪表盘)

具有不同纬度的仪表盘

### Histogram (直方图)

在一段时间范围内对数据进行采样

### Summary (摘要)

类似于Histogram，

## 参考

* [基于docker 搭建Prometheus+Grafana](https://www.cnblogs.com/xiao987334176/p/9930517.html)
