# ecryptfs

目录加密

## 安装

```shell
sudo apt install ecryptfs-utils
```

加载eCryptFS内核模块

```shell
sudo modprobe ecryptfs
```

初始化加密目录

```shell
ecryptfs-setup-private
```

加载目录

```shell
ecryptfs-mount-private ~/.Private ~/Private
```

拆卸目录

```shell
ecryptfs-umount-private ~/Private
```
