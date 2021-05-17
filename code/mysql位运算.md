# mysql位运算

假设有个表 test.bytes，表中有个字段b记录4种状态

* 0: 状态1
* 1: 状态2
* 2: 状态3
* 3: 状态4

4条记录，最终状态：

* id=1的记录状态1+状态2
* id=2的记录状态2+状态4
* id=3的记录状态1+状态2+状态3
* id=4的记录状态1+状态2+状态4

## 插入

```sql
-- id=1的记录状态1+状态2
INSERT INTO test.bytes (id, b) VALUES (1, 1<<0 | 1<<1);
-- id=2的记录状态2+状态4
INSERT INTO test.bytes (id, b) VALUES (2, 1<<1 | 1<<3);
-- id=3的记录状态1+状态2+状态3
INSERT INTO test.bytes (id, b) VALUES (3, 1<<0 | 1<<1 | 1<<2);
-- id=4的记录状态1+状态2
INSERT INTO test.bytes (id, b) VALUES (4, 1<<0 | 1<<1);

-- 1=3
-- 2=10
-- 3=7
-- 4=3
```

## 修改

```sql
-- 记录4增加状态3和状态4
UPDATE test.bytes SET b=b | (1<<2 | 1<<3) WHERE id=4;
-- b=15

-- 记录4删除状态4
UPDATE test.bytes SET b=b & ~(1<<3) WHERE id=4;
-- b=7

-- 记录4切换状态3和状态4
UPDATE test.bytes SET b=b ^ (1<<2 | 1<<3) WHERE id=4;
-- b=11
```

ps：切换状态是有这个状态就去掉，没有这个状态就增加

## 查询

```sql
-- 查询包含状态1的记录
SELECT * FROM test.bytes where b & 1<<0;
-- 结果: 1, 3, 4

-- 查询包含状态3或者状态4的记录
SELECT * FROM test.bytes where b & (1<<2|1<<3);
-- 结果: 2, 3, 4

-- 查询包含状态2和状态4的记录
SELECT * FROM test.bytes where b = b | 1<<1 | 1<<3;
-- 结果: 2, 4
```