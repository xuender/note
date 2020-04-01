# tail 内容显示

用于查看文件内容

## 选项

* -f 不停地去读最新的内容，这样有实时监视的效果 用Ctrl＋c来终止

## 技巧

* 程序标准输出重定向到文件，使用tail再同步输出到控制台

```bash
command > file & tail -f file
```

```flow
st=>start: Start
e=>end: End
op1=>operation: My Operation
sub1=>subroutine: My Subroutine
cond=>condition: Yes or No?
io=>inputoutput: catch something...
```
