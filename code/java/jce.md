# jce

这是因为某些国家的进口管制限制，JDK默认的加解密有一定的限制。
比如默认不允许 256 位密钥的 AES 加解密，解决方法就下载官方JCE无限制强度加密策略文件，覆盖即可。

## 安装

```shell
# yum -y curl unzip
curl -q -L -C - -b "oraclelicense=accept-securebackup-cookie" -o /tmp/jce_policy-8.zip \
       -O http://download.oracle.com/otn-pub/java/jce/8/jce_policy-8.zip \
    && unzip -oj -d ${JAVA_HOME}/jre/lib/security /tmp/jce_policy-8.zip \*/\*.jar \
    && rm /tmp/jce_policy-8.zip
```

## 参考

https://www.cnblogs.com/kancy/p/13276434.html
