package ta

import (
	"log"
	"reflect"
)

type comp struct {
	Base[bool]
	constant bool
	c        float64
}

func Comp(op func(float64, float64) bool, src Series[float64], val any) Series[bool] {
	var s comp
	var d []bool

	switch v := val.(type) {
	case int:
		d = make([]bool, 0, len(src.Data()))
		val := float64(v)
		for _, vv := range src.Data() {
			d = append(d, op(vv, val))
		}
	case float64:
		d = make([]bool, 0, len(src.Data()))
		for _, vv := range src.Data() {
			d = append(d, op(vv, v))
		}
	case *CONSTANT:
		d = make([]bool, 0, len(src.Data()))
		val := v.V(0)
		for _, vv := range src.Data() {
			d = append(d, op(vv, val))
		}
	case Series[float64]:
		var f, f1 []float64 = src.Data(), v.Data()
		l, pos := ShortestLenOfArray(f, f1)
		if pos == 0 {
			f1 = f1[len(f1)-l:]
		} else {
			f = f[len(f)-l:]
		}
		for i := 0; i < len(f); i++ {
			d = append(d, op(f[i], f1[i]))
		}
	default:
		log.Fatal("Comparison, not valid type ", reflect.TypeOf(v))
	}
	s.data = d
	return &s
}

// Smaller (src,v) => src < v
func Smaller(src Series[float64], v interface{}) Series[bool] {
	o := func(v1 float64, v2 float64) bool {
		return v1 < v2
	}
	return Comp(o, src, v)
}

// Greater (src,v) => src > v
func Greater(src Series[float64], v interface{}) Series[bool] {
	o := func(v1 float64, v2 float64) bool {
		return v1 > v2
	}
	return Comp(o, src, v)
}
