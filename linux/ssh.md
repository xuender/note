# ssh

## 免秘钥

```shell
ssh-copy-id -i ~/.ssh/id_rsa.pub 192.168.1.63
```

## 链接超时

客户端编辑 `/etc/ssh/ssh_config`

```txt
ServerAliveInterval 30
ServerAliveCountMax 3
```

客户端会在终端无操作之后 ServerAliveInterval 秒时请求服务器要求服务器响应，如果服务器在 ServerAliveCountMax 次之后都没有响应，则断开连接并退出。

## 建立远程端口转发隧道

将远程内网192.1.1.3的8080端口映射到本地

```shell
ssh -L 8080:192.1.1.3:8080 user@ip.ip.ip.ip -N 
```
