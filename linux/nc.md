# nc netcat

网络读取

## 启动监听

```shell
nc -l 9000
```

## 链接端口

```shell
nc localhost 9000
```

## 简单的聊天

1. 主机启动 `nc -l 9000`
1. 客户端链接 `nc xxx.xxx.xxx.xxx 9000`
1. 就可以聊天了

## 文件复制

1. 接收方启动 `nc -l 9000 > filename.out`
1. 发送方执行 `nc xxx.xxx.xxx.xxx 9000 < filename.in`

## 目录复制

1. 接收方启动 `nc -l 9000 |tar -C /目标目录 -zxf -`
1. 发送方执行 `tar -zcvf - 发送目录 | nc xxx.xxx.xxx.xxx 9000`

## 端口扫描

```shell
nc -v -z -w2 xxx.xxx.xxx.xxx 1-100
```

端口扫描不太好用，推荐使用[nmap](nmap.md)
