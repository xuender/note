# unar 解压

文件归档压缩格式众多，unar是一个可以解压多种归档类型的强力工具。  
已知可解压类型: zip, RAR, 7z, tar, gzip, bzip2, LZMA, XZ, CAB, MSI, NSIS, EXE, ISO, BIN 

和文件分割类型 Stuffit X, DiskDouble, Compact Pro, Packit, cpio, compress (.Z), ARJ, ARC, PAK, ACE, ZOO, LZH, ADF, DMS, LZX, PowerPacker, LBR, Squeeze, Crunch 以及其他不常用格式。

## 安装

```bash
sudo apt install unar
```

## 使用方法

```bash
unar [选项] 归档文件名
```

## 常用选项

* -o 输出目录
* -p 归档密码
* -r 如果文件、目录存在则重命名
