package ta

import "log"

type RSI struct {
	Base[float64]
}

func Rsi(src Series[float64], l int) *RSI {
	if l < 2 {
		log.Panicln("rsi invalid len", l)
	}
	r := new(RSI)
	f := src.Data()
	r.data = make([]float64, 0, len(f)-l)

	r.name = "RSI"
	gain, loss := avgGainLoss(f)
	avgGain, avgLoss := average(gain[:l]...), average(loss[:l]...)
	rs := avgGain / avgLoss
	r.data = append(r.data, 100-(100/(1+rs)))
	alpha := 1 / float64(l)
	alphaM := 1 - alpha
	for i := l; i < len(gain); i++ {
		avgGain = alpha*gain[i] + alphaM*avgGain
		avgLoss = alpha*loss[i] + alphaM*avgLoss
		r.data = append(r.data, 100-(100/(1+avgGain/avgLoss)))
	}

	//r.Gain = avgGain //Used for live
	//r.Loss = avgLoss //Used for live calculating
	return r
}

// Gets you the avg loss/gain for the rsi calculation
func avgGainLoss(f []float64) ([]float64, []float64) {
	gain := make([]float64, 0, len(f))
	loss := make([]float64, 0, len(f))
	gain = append(gain, 0)
	loss = append(loss, 0)
	var diff float64
	for i := 1; i < len(f); i++ {
		diff = f[i] - f[i-1]
		if diff >= 0 {
			gain = append(gain, diff)
			loss = append(loss, 0)
		} else {
			gain = append(gain, 0)
			loss = append(loss, -1*diff)
		}
	}
	return gain, loss
}
