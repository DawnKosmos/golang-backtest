package main

import (
	"fmt"
	"github.com/DawnKosmos/golang-backtest/internal/ta"
)

func main() {
	v := ta.Random(100, 20)
	r := ta.Rsi(v, 10)
	res := ta.Iff(ta.Smaller(r, v), r, ta.C(99))
	fmt.Println(res.Data())
}
