package ta

type IFF[T Type] struct {
	Base[T]
}

func Iff[T Type](con Series[bool], Then, Else Series[T]) *IFF[T] {
	s := new(IFF[T])
	// If Then or Else is of Type Constant. We have to generate the []Data
	thenVal, ok := any(Then).(*CONSTANT)
	if ok {
		thenVal.data = ConstArray(thenVal.Value, len(con.Data()))
	}
	elseVal, ok := any(Else).(*CONSTANT)
	if ok {
		elseVal.data = ConstArray(elseVal.Value, len(con.Data()))
	}

	c, t, e := con.Data(), Then.Data(), Else.Data()
	l := min(len(c), len(t), len(e))
	c = c[len(c)-l:]
	t = t[len(t)-l:]
	e = e[len(e)-l:]
	s.data = make([]T, l, l)
	for i, v := range c {
		if v {
			s.data[i] = t[i]
		} else {
			s.data[i] = e[i]
		}
	}

	return s
}
