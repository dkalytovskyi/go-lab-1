package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("buildVersion.txt")
	if err != nil {
		log.Fatal(err)
	}
	buildVersion := string(content)

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
