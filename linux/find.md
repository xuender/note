# find 文件查找

## find.sh

使用`find`和`grep`命令查找文件内容。

```bash
#! /bin/bash
# 文件内容搜索

if [ ! "$1" ] ;then
  echo "缺少查找的扩展文件名，例如: java"
fi
if [ ! "$2" ] ;then
  echo "缺少查找内容"
  echo
  echo "用法: find.sh 扩展文件名 查找内容"
  echo
  echo "例如: find.sh java DTO"
  exit
fi

endColour="\033[0m\e[0m"
redColour="\e[0;31m\033[1m"
purpleColour="\033[35m"

echo -e "查找扩展文件名: $purpleColour$1$endColour"
echo -e "查找内容: $redColour$2$endColour"
find ./ -name "*.$1" | xargs grep $2 --colour=auto
```
