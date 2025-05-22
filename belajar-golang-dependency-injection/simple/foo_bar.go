package simple

// ===========================//
// Struct Provider
// ===========================//
// membuat struct menjadi provider
type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

func NewBar() *Bar {
	return &Bar{}
}

type FooBar struct {
	*Foo
	*Bar
}
