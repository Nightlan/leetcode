package main

import "fmt"

func multiply(num1 string, num2 string) string {
	totalLen := len(num1) + len(num2)
	res := make([]byte, totalLen)
	for i := range res {
		res[i] = '0'
	}
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			tmp := (res[i+j+1] - '0') + (num1[i]-'0')*(num2[j]-'0')
			res[i+j+1] = tmp%10 + '0'
			res[i+j] += tmp / 10
		}
	}
	for i := 0; i < len(res); i++ {
		if res[i] != '0' {
			return string(res[i:])
		}
	}
	return "0"
}

func main() {
	num1 := "99965675765765765476"
	num2 := "9997677"
	fmt.Println(multiply(num1, num2))
	return
}
