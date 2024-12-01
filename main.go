package main

import (
	"log"

	"github.com/go-programming-tour-book/tour/cmd"
)

// 标准库flag的基本使用和长短选项
// func main() {
// 	var name string
// 	flag.StringVar(&name, "name", "Go语言", "帮助信息")
// 	flag.StringVar(&name, "n", "Go语言", "帮助信息")
// 	flag.Parse()

// 	log.Printf("name: %s", name)
// }

// 子命令使用
// var name string

// func main() {
// 	flag.Parse()
// 	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
// 	goCmd.StringVar(&name, "name", "go语言", "帮助信息")

// 	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
// 	phpCmd.StringVar(&name, "n", "php语言", "帮助信息")

// 	args := flag.Args()
// 	if len(args) <= 0 {
// 		return
// 	}

// 	switch args[0] {
// 	case "go":
// 		_ = goCmd.Parse(args[1:])
// 	case "php":
// 		_ = phpCmd.Parse(args[1:])
// 	}

// 	log.Printf("name: %s", name)
// }

// 自定义定义参数类型
// type Value interface {
// 	String() string
// 	Set(string) error
// }

// type Name string

// func (i *Name) String() string {
// 	return fmt.Sprint(*i)
// }

// func (i *Name) Set(value string) error {
// 	if len(*i) > 0 {
// 		return errors.New("name flag already set")
// 	}

// 	*i = Name("xiaowei:" + value)
// 	return nil
// }

// func main() {
// 	var name Name
// 	flag.Var(&name, "name", "帮助信息")
// 	flag.Parse()

// 	log.Printf("name: %s", name)
// }

// 单词格式转换

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
