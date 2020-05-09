# crontab

定时执行任务

## 编辑

```shell
crontab -e
```

## 提示

```crontab
* * * * * XDG_RUNTIME_DIR=/run/user/$(id -u) notify-send "提示" "信息"
```
