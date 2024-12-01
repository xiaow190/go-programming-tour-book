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