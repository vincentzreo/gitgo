package main

import (
	"fmt"
	colour "github.com/fatih/color"
	flag "github.com/ogier/pflag"
	"os"
	"strings"
)

var (
	user  string
	emoji string
)

func main() {
	// 运行 go run main.go
	// 建造 go build
	// 安装 go install
	flag.Parse()
	if flag.NFlag() == 0 {
		printUsage()
	}
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	fmt.Println("")
	for _, u := range users {
		result := getUesrs(u)
		colour.Cyan(`Username: %s`, result.Login)
		colour.Blue(`Name:			%s`, result.Name)
		colour.Green(`Email:			%s`, result.Email)
		colour.HiMagenta(`Bio:		%s`, result.Bio)
		fmt.Println("")
	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
