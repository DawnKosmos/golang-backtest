package live

import "github.com/DawnKosmos/golang-backtest/internal/ta"

type Series[T ta.Type] interface {
	V(index int) T
	Latest() T
	Data() []T
	Name() string
	SetName(name string)
	Updater
}

type Updater interface {
	OnTick(newCandle bool) // Calculation that is Done on a new tick
	UpdateGroup() UpdateGroup
	SetLimit(i int) //Sets the Limit that needs to be allocated for the indicator to work
	ExecuteLimit()  //Gets called once
}

type UpdateGroup interface {
	Add(updater Updater)
}

type Base[T ta.Type] struct {
	src    Series[T]
	recent T
	d      Dater[T]
	name   string
	limit  int
	ug     UpdateGroup
}

func (b *Base[T]) SetBase(u Updater, src Series[T]) {
	b.ug = src.UpdateGroup()
	b.src = src
	b.ug.Add(u)
}

func (b *Base[T]) UpdateGroup() UpdateGroup {
	return b.ug
}

func (b *Base[T]) Name() string {
	return b.name
}

func (b *Base[T]) SetName(name string) {
	b.name = name
}

func (b *Base[T]) V(index int) T {
	return b.d.V(index)
}

func (b *Base[T]) Data() []T {
	return b.d.Data()
}

func (b *Base[T]) SetLimit(limit int) {
	if limit > b.limit {
		b.limit = limit
	}
}

func (b *Base[T]) ExecuteLimit() {
	b.d.SetLimit(b.limit)
}

func (b *Base[T]) Latest() T {
	return b.d.V(0)
}
