# man 文档

## 中文文档

```bash
sudo apt install manpages-zh
vi /etc/manpath.config
:1,$s#/usr/share/man#/usr/share/man/zh_CN#g
```

将所有的/usr/share/man替换为/usr/share/man/zh_CN
