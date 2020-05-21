# delete

## 按照时间删除缺少时间索引的表

如果想按照时间删除`id`是自增主键的表，但缺少时间索引。

执行以下查询语句，看`created`是否在删除范围

```sql
SELECT id, created FROM t_table ORDER BY id LIMIT 1000000, 1;
```

然后反复执行以下语句，每次执行完检查`created`是否在删除范围

```sql
DELETE FROM t_table ORDER BY id LIMIT 1000000;
SELECT id,created FROM t_table ORDER BY id LIMIT 1000000, 1;
```

可根据表大小调整`LIMIT`尺寸
