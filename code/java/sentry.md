# sentry 异常追踪

## maven

```xml
<dependency>
  <groupId>io.sentry</groupId>
  <artifactId>sentry-spring</artifactId>
  <version>1.7.30</version>
</dependency>
<dependency>
  <groupId>io.sentry</groupId>
  <artifactId>sentry-logback</artifactId>
  <version>1.7.30</version>
</dependency>
```

## logback

```xml
<appender name="Sentry" class="io.sentry.logback.SentryAppender">
  <filter class="ch.qos.logback.classic.filter.ThresholdFilter">
    <level>ERROR</level>
  </filter>
</appender>
<root level="info">
  <appender-ref ref="console"/>
  <appender-ref ref="file"/>
  <appender-ref ref="Sentry"/>
</root>
```

## 环境变量增加

```shell
SENTRY_DSN
```
