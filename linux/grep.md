# grep

文件内容查找

## 查找JSON值

```shell
grep '138xxxxx' /tmp/a.txt | grep waybill_code | sed -e 's/^.*"waybill_code":"\([^"]*\)".*$/\1/'
```
