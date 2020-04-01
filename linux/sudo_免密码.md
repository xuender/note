# sudo 免密码

经常执行远程命令

```shell
ssh domain.cn "~/bin/command.sh"
```

如果`command.sh`中有`sudo`命令就无法执行，返回

```shell
sudo: no tty present and no askpass program specified
```

## 免密方法

```shell
sudo chmod u+w /etc/sudoers
```

```shell
sudo vi /etc/sudoers
```

将`%sudo   ALL=(ALL:ALL) ALL` 改成 `%sudo   ALL=(ALL:ALL) NOPASSWD:ALL` 即可。

最后

```shell
sudo chmod u-w /etc/sudoers
```

## 异常

```shell
sudo: /etc/sudoers 可被任何人写
sudo: 没有找到有效的 sudoers 资源，退出
sudo: 无法初始化策略插件
```

修改权限

```shell
pkexec chmod 0440 /etc/sudoers
```

然后重来
