# logrotate

日志拆分

`/etc/logrotate.conf` 主配置文件
日志配置目录 `/etc/logrotate.d`

## 常用命令

手动执行某日志拆分命令

```shell
logrotate /etc/logrotate.d/log-file
```

调试某日志拆分命令，并不执行

```shell
logrotate -d /etc/logrotate.d/log-file
```

## 配置文件

* su banboo banboo 是 su `user` `group`
* missingok表示如果找不到log文件也没OK
* compress 通过gzip 压缩转储以后的日志
* nocompress 不需要压缩时，用这个参数
* rotate 31 表示保留31天的备份文件
* notifempty 表示如果log文件是空的，就不进行rotate
* copytruncate 表示先复制log文件的内容，然后再清空
* dateext表示备份的日志文件后缀格式为YYYYMMDD

## 例子

```config
/var/tmp/crm.log {
    su root ender
    daily
    rotate 10
    compress
    delaycompress
    missingok
    notifempty
    copytruncate
    dateext
    create 644 ender ender
    postrotate
        su - ender -c "/home/ender/work/crm/crm stop"
        su - ender -c "cd /home/ender/work/crm && ./crm start"
    endscript
}
```
