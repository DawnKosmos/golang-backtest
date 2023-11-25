package test

import (
	"github.com/DawnKosmos/golang-backtest/internal/live"
	"github.com/DawnKosmos/golang-backtest/internal/ta"
)

type taFunc[T ta.Type] func(ta.Series[T], int) ta.Series[T]
type liveFunc[T ta.Type] func(live.Series[T], int) live.Series[T]

func Compare[T ta.Type](srcA ta.Series[T], srcB live.Series[T]) {

}
