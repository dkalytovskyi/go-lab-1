package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(buildVersion)

	args := os.Args[1:]
	args_str := strings.Join(args, " ")
	res, err := PrefixToPostfix(args_str)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err.Error())
	}

}
