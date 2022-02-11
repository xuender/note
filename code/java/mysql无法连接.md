# mysql 无法连接

无法连接

Caused by: javax.net.ssl.SSLHandshakeException: No appropriate protocol (protocol is disabled or cipher suites are inappropriate)


## 修改

修改 jre/lib/security/java.security

将 jdk.tls.disabledAlgorithms 注释掉

```text
# jdk.tls.disabledAlgorithms=SSLv3, TLSv1, TLSv1.1, RC4, DES, MD5withRSA, \
#    DH keySize < 1024, EC keySize < 224, 3DES_EDE_CBC, anon, NULL, \
#    include jdk.disabled.namedCurves
```

## 堆栈

```text
org.apache.ibatis.exceptions.PersistenceException: 
### Error querying database.  Cause: com.hexin.light.spring.jdbc.CannotGetJdbcConnectionException: Could not get JDBC Connection
### The error may exist in file [/home/ender/work/21/eis/target/classes/mybatis/waybillLogMapper.xml]
### The error may involve com.hexin.eis.waybill.WaybillLogDAO.query
### The error occurred while executing a query
### Cause: com.hexin.light.spring.jdbc.CannotGetJdbcConnectionException: Could not get JDBC Connection
	at org.apache.ibatis.exceptions.ExceptionFactory.wrapException(ExceptionFactory.java:30)
	at org.apache.ibatis.session.defaults.DefaultSqlSession.selectList(DefaultSqlSession.java:149)
	at org.apache.ibatis.session.defaults.DefaultSqlSession.selectList(DefaultSqlSession.java:140)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:498)
	at com.hexin.light.mybatis.spring.SqlSessionTemplate$SqlSessionInterceptor.invoke(SqlSessionTemplate.java:394)
	at com.sun.proxy.$Proxy46.selectList(Unknown Source)
	at com.hexin.light.mybatis.spring.SqlSessionTemplate.selectList(SqlSessionTemplate.java:194)
	at com.hexin.eis.waybill.WaybillLogDAO.query(WaybillLogDAO.java:78)
	at com.hexin.eis.waybill.impl.WaybillServiceImpl.getBills(WaybillServiceImpl.java:310)
	at com.hexin.eis.waybill.WaybillCtrl.getBills(WaybillCtrl.java:86)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:498)
	at com.hexin.light.mvc.ctrl.AbstractController.doPost(AbstractController.java:147)
	at javax.servlet.http.HttpServlet.service(HttpServlet.java:707)
	at javax.servlet.http.HttpServlet.service(HttpServlet.java:790)
	at com.google.inject.servlet.ServletDefinition.doServiceImpl(ServletDefinition.java:290)
	at com.google.inject.servlet.ServletDefinition.doService(ServletDefinition.java:280)
	at com.google.inject.servlet.ServletDefinition.service(ServletDefinition.java:184)
	at com.google.inject.servlet.ManagedServletPipeline.service(ManagedServletPipeline.java:89)
	at com.google.inject.servlet.FilterChainInvocation.doFilter(FilterChainInvocation.java:85)
	at com.google.inject.servlet.ManagedFilterPipeline.dispatch(ManagedFilterPipeline.java:121)
	at com.google.inject.servlet.GuiceFilter.doFilter(GuiceFilter.java:133)
	at org.eclipse.jetty.servlet.FilterHolder.doFilter(FilterHolder.java:193)
	at org.eclipse.jetty.servlet.ServletHandler$Chain.doFilter(ServletHandler.java:1601)
	at com.hexin.light.web.CharsetFilter.doFilter(CharsetFilter.java:29)
	at org.eclipse.jetty.servlet.FilterHolder.doFilter(FilterHolder.java:201)
	at org.eclipse.jetty.servlet.ServletHandler$Chain.doFilter(ServletHandler.java:1601)
	at org.eclipse.jetty.servlet.ServletHandler.doHandle(ServletHandler.java:548)
	at org.eclipse.jetty.server.handler.ScopedHandler.handle(ScopedHandler.java:143)
	at org.eclipse.jetty.security.SecurityHandler.handle(SecurityHandler.java:602)
	at org.eclipse.jetty.server.handler.HandlerWrapper.handle(HandlerWrapper.java:127)
	at org.eclipse.jetty.server.handler.ScopedHandler.nextHandle(ScopedHandler.java:235)
	at org.eclipse.jetty.server.session.SessionHandler.doHandle(SessionHandler.java:1624)
	at org.eclipse.jetty.server.handler.ScopedHandler.nextHandle(ScopedHandler.java:233)
	at org.eclipse.jetty.server.handler.ContextHandler.doHandle(ContextHandler.java:1435)
	at org.eclipse.jetty.server.handler.ScopedHandler.nextScope(ScopedHandler.java:188)
	at org.eclipse.jetty.servlet.ServletHandler.doScope(ServletHandler.java:501)
	at org.eclipse.jetty.server.session.SessionHandler.doScope(SessionHandler.java:1594)
	at org.eclipse.jetty.server.handler.ScopedHandler.nextScope(ScopedHandler.java:186)
	at org.eclipse.jetty.server.handler.ContextHandler.doScope(ContextHandler.java:1350)
	at org.eclipse.jetty.server.handler.ScopedHandler.handle(ScopedHandler.java:141)
	at org.eclipse.jetty.server.handler.ContextHandlerCollection.handle(ContextHandlerCollection.java:191)
	at org.eclipse.jetty.server.handler.HandlerCollection.handle(HandlerCollection.java:146)
	at org.eclipse.jetty.server.handler.HandlerWrapper.handle(HandlerWrapper.java:127)
	at org.eclipse.jetty.server.Server.handle(Server.java:516)
	at org.eclipse.jetty.server.HttpChannel.lambda$handle$1(HttpChannel.java:388)
	at org.eclipse.jetty.server.HttpChannel.dispatch(HttpChannel.java:633)
	at org.eclipse.jetty.server.HttpChannel.handle(HttpChannel.java:380)
	at org.eclipse.jetty.server.HttpConnection.onFillable(HttpConnection.java:273)
	at org.eclipse.jetty.io.AbstractConnection$ReadCallback.succeeded(AbstractConnection.java:311)
	at org.eclipse.jetty.io.FillInterest.fillable(FillInterest.java:105)
	at org.eclipse.jetty.io.ChannelEndPoint$1.run(ChannelEndPoint.java:104)
	at org.eclipse.jetty.util.thread.QueuedThreadPool.runJob(QueuedThreadPool.java:773)
	at org.eclipse.jetty.util.thread.QueuedThreadPool$Runner.run(QueuedThreadPool.java:905)
	at java.lang.Thread.run(Thread.java:748)
Caused by: com.hexin.light.spring.jdbc.CannotGetJdbcConnectionException: Could not get JDBC Connection
	at com.hexin.light.spring.jdbc.datasource.DataSourceUtils.getConnection(DataSourceUtils.java:85)
	at com.hexin.light.mybatis.spring.transaction.SpringManagedTransaction.openConnection(SpringManagedTransaction.java:79)
	at com.hexin.light.mybatis.spring.transaction.SpringManagedTransaction.getConnection(SpringManagedTransaction.java:65)
	at org.apache.ibatis.executor.BaseExecutor.getConnection(BaseExecutor.java:337)
	at org.apache.ibatis.executor.SimpleExecutor.prepareStatement(SimpleExecutor.java:86)
	at org.apache.ibatis.executor.SimpleExecutor.doQuery(SimpleExecutor.java:62)
	at org.apache.ibatis.executor.BaseExecutor.queryFromDatabase(BaseExecutor.java:325)
	at org.apache.ibatis.executor.BaseExecutor.query(BaseExecutor.java:156)
	at org.apache.ibatis.executor.CachingExecutor.query(CachingExecutor.java:109)
	at org.apache.ibatis.executor.CachingExecutor.query(CachingExecutor.java:89)
	at org.apache.ibatis.session.defaults.DefaultSqlSession.selectList(DefaultSqlSession.java:147)
	... 58 common frames omitted
Caused by: com.mysql.cj.jdbc.exceptions.CommunicationsException: Communications link failure

The last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.
	at com.mysql.cj.jdbc.exceptions.SQLError.createCommunicationsException(SQLError.java:174)
	at com.mysql.cj.jdbc.exceptions.SQLExceptionsMapping.translateException(SQLExceptionsMapping.java:64)
	at com.mysql.cj.jdbc.ConnectionImpl.createNewIO(ConnectionImpl.java:828)
	at com.mysql.cj.jdbc.ConnectionImpl.<init>(ConnectionImpl.java:448)
	at com.mysql.cj.jdbc.ConnectionImpl.getInstance(ConnectionImpl.java:241)
	at com.mysql.cj.jdbc.NonRegisteringDriver.connect(NonRegisteringDriver.java:198)
	at com.zaxxer.hikari.util.DriverDataSource.getConnection(DriverDataSource.java:138)
	at com.zaxxer.hikari.pool.PoolBase.newConnection(PoolBase.java:358)
	at com.zaxxer.hikari.pool.PoolBase.newPoolEntry(PoolBase.java:206)
	at com.zaxxer.hikari.pool.HikariPool.createPoolEntry(HikariPool.java:477)
	at com.zaxxer.hikari.pool.HikariPool.checkFailFast(HikariPool.java:560)
	at com.zaxxer.hikari.pool.HikariPool.<init>(HikariPool.java:115)
	at com.zaxxer.hikari.HikariDataSource.getConnection(HikariDataSource.java:112)
	at com.hexin.light.spring.jdbc.datasource.DataSourceUtils.doGetConnection(DataSourceUtils.java:116)
	at com.hexin.light.spring.jdbc.datasource.DataSourceUtils.getConnection(DataSourceUtils.java:82)
	... 68 common frames omitted
Caused by: com.mysql.cj.exceptions.CJCommunicationsException: Communications link failure

The last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.
	at sun.reflect.NativeConstructorAccessorImpl.newInstance0(Native Method)
	at sun.reflect.NativeConstructorAccessorImpl.newInstance(NativeConstructorAccessorImpl.java:62)
	at sun.reflect.DelegatingConstructorAccessorImpl.newInstance(DelegatingConstructorAccessorImpl.java:45)
	at java.lang.reflect.Constructor.newInstance(Constructor.java:423)
	at com.mysql.cj.exceptions.ExceptionFactory.createException(ExceptionFactory.java:61)
	at com.mysql.cj.exceptions.ExceptionFactory.createException(ExceptionFactory.java:105)
	at com.mysql.cj.exceptions.ExceptionFactory.createException(ExceptionFactory.java:151)
	at com.mysql.cj.exceptions.ExceptionFactory.createCommunicationsException(ExceptionFactory.java:167)
	at com.mysql.cj.protocol.a.NativeProtocol.negotiateSSLConnection(NativeProtocol.java:317)
	at com.mysql.cj.protocol.a.NativeAuthenticationProvider.connect(NativeAuthenticationProvider.java:202)
	at com.mysql.cj.protocol.a.NativeProtocol.connect(NativeProtocol.java:1352)
	at com.mysql.cj.NativeSession.connect(NativeSession.java:132)
	at com.mysql.cj.jdbc.ConnectionImpl.connectOneTryOnly(ConnectionImpl.java:948)
	at com.mysql.cj.jdbc.ConnectionImpl.createNewIO(ConnectionImpl.java:818)
	... 80 common frames omitted
Caused by: javax.net.ssl.SSLHandshakeException: No appropriate protocol (protocol is disabled or cipher suites are inappropriate)
	at sun.security.ssl.HandshakeContext.<init>(HandshakeContext.java:171)
	at sun.security.ssl.ClientHandshakeContext.<init>(ClientHandshakeContext.java:98)
	at sun.security.ssl.TransportContext.kickstart(TransportContext.java:220)
	at sun.security.ssl.SSLSocketImpl.startHandshake(SSLSocketImpl.java:428)
	at com.mysql.cj.protocol.ExportControlled.performTlsHandshake(ExportControlled.java:320)
	at com.mysql.cj.protocol.StandardSocketFactory.performTlsHandshake(StandardSocketFactory.java:194)
	at com.mysql.cj.protocol.a.NativeSocketConnection.performTlsHandshake(NativeSocketConnection.java:101)
	at com.mysql.cj.protocol.a.NativeProtocol.negotiateSSLConnection(NativeProtocol.java:308)
	... 85 common frames omitted
```
