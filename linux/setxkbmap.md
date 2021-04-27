# setxkbmap

修改加盘配置

## CapsLock 改成 ctrl

```shell
setxkbmap -option caps:ctrl_modifier
```

## 启动修改

/etc/default/keyboard

`XKBOPTIONS="caps:ctrl_modifier"`
