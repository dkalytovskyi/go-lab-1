package main

import (
	"strings"
	"errors"
	"strconv"
)

func IsOperator(c string) bool {
	return strings.ContainsAny(string(c), "+ & - & * & / & ^")
}

func IsValidOperand(c string) bool {
	 _, err := strconv.ParseFloat(c, 64)
   return err == nil
}

func InputCheck(c []string) (bool, error) {
	length := len(c)

	operands_num := 0
	operators_num := 0 

	if(length == 0){
		return false, errors.New("Error: empty line entered!")
	}

	if(length < 3){
		return false, errors.New("Error: not enough operands!")
	}

	for i := 0; i < length; i++ {
		if(IsOperator(string(c[i]))){
			operators_num += 1
		} else if (IsValidOperand(string(c[i]))){
			operands_num += 1
		} else {
			return false, errors.New("Error: operand is not valid!")
		}
	}

	if (operators_num + 1 != operands_num){
		return false, errors.New("Error: input is not valid")
	}

	return true,nil
}

func PrefixToPostfix(s string) (string, error) {
	var stack []string

	s_elem := strings.Fields(s)

	res, err := InputCheck(s_elem)
	if (res == false){
		return "",err
	}

	length := len(s_elem)

	for i := length-1; i >= 0; i-- {
		n := len(stack) - 1

		if(IsOperator(s_elem[i])){
			op1 := stack[n]
			op2 := stack[n-1]

			str := op1+" "+op2+" "+string(s_elem[i])

			stack = stack[:n-1]
			stack = append(stack, str)
			
		} else {

			stack = append(stack, string(s_elem[i]))
		}
	}
	return stack[0], nil
}
