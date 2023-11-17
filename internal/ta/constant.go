package ta

import "math/rand"

type CONSTANT struct {
	Base[float64]
	Value float64
}

func C(v float64) *CONSTANT {
	return &CONSTANT{Value: v}
}

func constantWithLen(v float64, l int) *CONSTANT {
	c := &CONSTANT{Value: v}
	c.data = make([]float64, l, l)
	for i := 0; i < l; i++ {
		c.data[i] = v
	}
	return c
}

func Constant(Value float64) *CONSTANT {
	return &CONSTANT{Value: Value}
}

func (C *CONSTANT) V(_ int) float64 {
	return C.Value
}

func (C *CONSTANT) Name() string {
	return "Constant"
}

func (C *CONSTANT) SetName(_ string) {
	return
}

func Random(max int32, l int) *CONSTANT {
	c := &CONSTANT{}
	c.data = make([]float64, l, l)
	for i := 0; i < l; i++ {
		c.data[i] = float64(rand.Int31n(max))
	}
	return c
}
