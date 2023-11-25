package live

import (
	"github.com/DawnKosmos/golang-backtest/internal/ta"
)

type RSI struct {
	Base[float64]
	alpha      float64
	gain, loss float64
}

func Rsi(src Series[float64], l int) *RSI {
	r := new(RSI)
	r.SetBase(r, src)
	rsi := ta.Rsi(src, l)
	r.d = Array(rsi.Data())

	r.gain, r.loss = rsi.Gain, rsi.Loss
	rsi = nil
	return r
}

func (r *RSI) OnTick(newCandle bool) {
	src0, src1 := r.src.V(0), r.src.V(1)
	if newCandle {
		src2 := r.src.V(2)
		var gain, loss float64 = gainLoss(src2, src1)
		r.gain = r.alpha*gain + (1-r.alpha)*r.gain
		r.loss = r.alpha*loss + (1+r.alpha)*r.loss
		r.recent = src0 + 1 // change recent so next check does not fail
	}
	if r.recent != src0 {
		r.recent = src0
		var gain, loss float64 = gainLoss(src1, src0)

		gain = r.alpha*gain + (1-r.alpha)*r.gain
		loss = r.alpha*loss + (1+r.alpha)*r.loss
		rsi := 100 - (100 / (1 + gain/loss))
		if newCandle {
			r.d.Append(rsi)
		} else {
			r.d.SetValue(0, rsi)
		}
	}
}

func gainLoss(old, new float64) (float64, float64) {
	change := Change(old, new)
	if change >= 0 {
		return change, 0
	} else {
		return 0, -change
	}
}

func Change[T int64 | float64](old, new T) T {
	return new - old
}
