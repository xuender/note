# 阿里云

## 时序数据库

```sql
-- 查询首页访问情况
* | select promql_query_range('COUNT_OF_WODA_HOME_PERIOD') from metrics

-- 过滤查询ERP相关的打印量
* | select promql_query_range('SUM_OF_waybillAmount_PRINT_PERIOD{app="woda_erp"}') from metrics

-- 显示每小时打单曲线
* | select promql_query_range('irate(SUM_OF_waybillAmount_PRINT_PERIOD{app="woda_erp"}[1h])') from metrics
```
