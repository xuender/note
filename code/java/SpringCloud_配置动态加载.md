# SpringCloud 配置动态加载

## 修改配置

application.yml 增加配置，启动 `actuator`。

```yaml
management:
  security:
    enabled: false
  endpoints:
    web:
      exposure:
        include: bus-refresh # 通过消息总线，刷新配置
  endpoint:
    bus-refresh:
      enabled: true
      sensitive: false # 密码验证
```

## 重新加载

执行命令

```shell
curl http://localhost:9302/actuator/bus-refresh -H "Authorization:Bearer JWT token..." -X POST
```
