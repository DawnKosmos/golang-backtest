package ta

type calc interface {
	int | int64 | float64 | float32 | int32
}

func average[T calc](f ...T) T {
	return summe[T](f...) / (T)(len(f))
}

func summe[T calc](f ...T) T {
	var avg T = 0
	for _, v := range f {
		avg += v
	}
	return avg
}

// ShortestLenOfArray Return the lenght of the shortest array
func ShortestLenOfArray[T any](f ...[]T) (int, int) {
	var l int = len(f[0])
	var position int = 0
	for i := 1; i < len(f); i++ {
		if len(f[i]) < l {
			l = len(f[i])
			position = i
		}
	}
	return l, position
}

func ConstArray(v float64, l int) []float64 {
	out := make([]float64, l, l)
	for i := 0; i < l; i++ {
		out[i] = v
	}
	return out
}
