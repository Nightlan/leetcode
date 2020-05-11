package main

import "fmt"

func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := 1.0
	tmpVal := x
	if n < 0 {
		tmpVal = 1 / tmpVal
		n = -n
	}
	for i := 0; i < n; i++ {
		res = res * tmpVal
	}
	return res
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := n / 2
	tmpRes := quickPow(x, half)
	if n%2 != 0 {
		return tmpRes * tmpRes * x
	}
	return tmpRes * tmpRes
}

func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickPow(x, n)
	}
	return 1 / quickPow(x, -n)
}

func main() {
	fmt.Println(myPow(2.0, -2))
}
