# CLI 命令行实用程序开发基础

## 1.概述

CLI（Command Line Interface）实用程序是Linux下应用开发的基础。正确的编写命令行程序让应用与操作系统融为一体，通过shell或script使得应用获得最大的灵活性与开发效率。

本实验要求开发一个go语言版本的selpg程序.

## 2.基础知识

- 使用os包,能调用os包所提供的接口,如文件读写,异常返回等
- 使用flag包,解析命令行参数

## 3.设计说明

该程序基于原版C语言的selpg

## 4.实现细节

导入相应的包

```go
import (
	"bufio"
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)
```

命令行参数结构体:

```go
type selpgArgs struct {
	startPage int
	endPage   int

	pageLen  int
	pageType bool

	printDest string

	inFileName string
}
```

解析参数:

```go
flag.IntVarP(&sa.startPage, "s", "s", -1, "startPage")
...
flag.Parse()
```

处理异常:

```go
if sa.startPage < 1 || sa.startPage > sa.endPage {
    // flag.Usage("wrong page number")
    flag.Usage()
    os.Exit(1)
}
...
```

selpg主要逻辑:

如果按特定行数换行,如果没有指定换页行数,则默认为72行,否则就按照指定行数换页

如果按换页符分页,读到换页符就换页

```go
if sa.pageType == false {
    ...
} else {
    ...
}
```

## 5.测试结果

脚本或者C语言生成in.txt,每一行为当前的行号

```
1
2
3
4
...
```



1. 把in.txt的第 1 页写至标准输出（也就是屏幕），因为这里没有重定向或管道。

   ```bash
   D:\goProject\src\github.com\runfeng1998\selpg>selpg -s 1 -e 1 in.txt
   1
   2
   3
   ...
   72
   ```

2. selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“in.txt”而不是显式命名的文件名参数。

   ```bash
   D:\goProject\src\github.com\runfeng1998\selpg>selpg -s 1 -e 1 < in.txt
   1
   2
   3
   ...
   72
   ```

3. “other_command”的标准输出被 shell／内核重定向至 selpg 的标准输入。

   ```bash
   D:\goProject\src\github.com\runfeng1998\selpg>type in.txt | selpg -s 10 -e 20
   649
   650
   651
   652
   653
   654
   ...
   1440
   ```

4. selpg 将第 10 页到第 20 页写至标准输出；标准输出被 shell／内核重定向至“out.txt”。

   ```bash
   D:\goProject\src\github.com\runfeng1998\selpg>selpg
   -s 10 -e 20 in.txt>out.txt
   D:\goProject\src\github.com\runfeng1998\type out.txt
   649
   650
   651
   652
   653
   654
   ...
   1440
   ```

5. selpg 将第 10 页到第 20 页写至标准输出（屏幕）；所有的错误消息被 shell／内核重定向至“error.txt”。

   ```bash
   D:\goProject\src\github.com\runfeng1998\selpg>selpg
   -s 10 -e 20 in.txt 2>error.txt
   649
   650
   651
   652
   653
   654
   ...
   1440
   ```

   error.txt无内容





