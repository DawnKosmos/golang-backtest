package ta

type Type interface {
	~bool | ~float64
}

type Series[T Type] interface {
	V(index int) T       // return the last minus index value. So V(0) returns the latest Candle
	Data() []T           // returns all known data of the indicator
	Name() string        // Name is needed for identification
	SetName(name string) // SetName also used for identification
}

type Base[T Type] struct {
	data []T
	name string
}

func (b *Base[T]) Data() []T {
	return b.data
}

func (b *Base[T]) V(index int) T {
	return b.data[len(b.data)-1-index]
}
func (b *Base[T]) Name() string {
	return b.name
}

func (b *Base[T]) SetName(name string) {
	b.name = name
}
