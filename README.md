# go-programming-tour-book
打造属于自己的工具集

## 单词格式转换 -- word
> 基于但三方开源库Cobra和标准库strings、unicode 实现了多种模式的单词转换功能

### 验证
```go

go run main.go help  word

该子命令支持各种单词格式转换，模式如下：
1: 全部单词转为大写
2: 全部单词转为小写
3: 下画线单词转为大写驼峰单词
4: 下画线单词转为小写驼峰单词
5: 驼峰单词转为下画线单词

Usage:
   word [flags]

Flags:
  -h, --help         help for word
  -m, --mode int8    请输入单词转换模式
  -s, --str string   请输入单词内容

```
```go
go run main.go word -s=xiaowei -m=1
输出结果: XIAOWEI

go run main.go word -s=XIAOWEI -m=2
输出结果: xiaowei

go run main.go word -s=xiaowei -m=3
输出结果: Xiaowei

go run main.go word -s=XIAOWEI -m=4
输出结果: xIAOWEI

go run main.go word -s=XiaoWei -m=5
输出结果: xiao_wei

```

## 便捷的时间工具， 

> 格式化时间

### 验证 

```go
go run main.go time  duration -h

时间格式处理

Usage:
   time [flags]
   time [command]

Available Commands:
  calc        计算所需时间
  now         获取当前时间

Flags:
  -h, --help   help for time

Use " time [command] --help" for more information about a command.


// 查看命令介绍
go run main.go time now -h
获取当前时间

Usage:
   time now [flags]

Flags:
  -h, --help   help for now

go run main.go time calc -h
计算所需时间

Usage:
   time calc [flags]

Flags:
  -c, --calculate string   需要计算的时间，有效单位为时间戳或已格式化后的时间
  -d, --duration string    持续时间，有效时间单位为"ns", "us"(or "us"), "ms", "s", "m", "h"
  -h, --help               help for calc

// 获取当前时间和时间戳
go run main.go time  now
输出结果: 2024-12-02 16:37:53, 1733128673

// 获取指定时间往后5分钟的时间和时间戳
>go run main.go time calc -c="2029-09-04 12:02:33" -d=5m
输出结果: 2029-09-04 12:07:33, 1883218053

// 获取指定时间往前2小时的时间和时间戳
go run main.go time calc -c="2029-09-04 12:02:33" -d=-2h
输出结果: 2029-09-04 10:02:33, 1883210553



```
