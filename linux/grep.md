# grep

文件内容查找

## 查找JSON值

```shell
grep '138xxxxx' /tmp/a.txt | grep waybill_code | sed -e 's/^.*"waybill_code":"\([^"]*\)".*$/\1/'
```

## 正则表达式

-E 扩展正则表达式
-o 只输出匹配到的内容

```shell
grep -E -o '\W[a-z]{2,3}:' trace.log > express.txt
```
