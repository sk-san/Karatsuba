package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(_Karatsuba("123456789", "987654321"))
}

func Karatsuba(x, y *big.Int) *big.Int {
	if x.Cmp(big.NewInt(10)) == -1 || y.Cmp(big.NewInt(10)) == -1 {
		result := new(big.Int)
		return result.Mul(x, y)
	} else {
		n := max(len(x.String()), len(y.String()))
		half := n / 2

		a := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(half)), nil) // 123456 -> 123000
		a.Div(x, a)
		b := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(half)), nil) // 123456 -> 456
		b.Mod(x, b)
		c := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(half)), nil)
		c.Div(y, c)
		d := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(half)), nil)
		d.Mod(y, d)

		// Step1
		ac := Karatsuba(a, c)
		// Step2
		bd := Karatsuba(b, d)
		// Step3 and 4
		a_b := new(big.Int).Add(a, b)
		c_d := new(big.Int).Add(c, d)
		ac_plus_bd := new(big.Int).Add(ac, bd)

		ac_bd := Karatsuba(a_b, c_d)
		ac_bd.Sub(ac_bd, ac_plus_bd)

		//Step5
		//return int(math.Pow10(half*2))*ac + int(math.Pow10(half))*ac_bd + bd
		half_1 := big.NewInt(int64(half) * 2)
		half_2 := big.NewInt(int64(half))

		result_left := new(big.Int).Exp(big.NewInt(10), half_1, nil)
		result_right := new(big.Int).Exp(big.NewInt(10), half_2, nil)

		result_left.Mul(result_left, ac)
		result_right.Mul(result_right, ac_bd)

		output := new(big.Int)
		output.Add(output, result_left)
		output.Add(output, result_right)
		output.Add(output, bd)

		return output
	}
}

func _Karatsuba(x, y string) int {
	nx, _ := strconv.Atoi(x)
	ny, _ := strconv.Atoi(y)
	if nx < 10 || ny < 10 {
		return nx * ny
	}

	m := max(len(x), len(y))
	half := m / 2

	xpos := len(x) - half
	ypos := len(y) - half
	a, c := x[:xpos], y[:ypos]
	b, d := x[xpos:], y[ypos:]

	ac := _Karatsuba(a, c)
	bd := _Karatsuba(b, d)
	ad := _Karatsuba(a, d)
	bc := _Karatsuba(b, c)

	adbc := ad + bc

	return ac*int(math.Pow10(half*2)) + adbc*int(math.Pow10(half)) + bd
}
