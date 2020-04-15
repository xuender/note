# netstat 端口状态

## 参数

* `-t` tcp
* `-u` udp
* `-l` 显示监听端口
* `-p` 显示占用的程序

## 端口占用

```shell
sudo netstat -tunlp | grep 80
```
