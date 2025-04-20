//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

// ===========================//
// Injector
// ===========================//
func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService, // memberitahukan wire untuk provider mana yang di inject
	)

	return nil, nil
}

func InitializeDatabase() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)

	return nil
}

//============================//
//Provider Set / Grouping
//============================//
//Menggabungkan beberapa provider dalam satu variable

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(
		fooSet, barSet, NewFooBarService,
	)

	return nil
}

// ============================//
// Binding Interface
// ============================//
func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

// ===========================//
// Struct Provider
// ===========================//
func InitializeFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar")) // (*) for all field
	return nil
}

// ===========================//
// Binding Values
// ===========================//
var fooValue = &Foo{}
var barValue = &Bar{}

func InitializeFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

// yang sudah ada valuenya
func InitializeReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// ==========================//
// Struct Field Provider
// ==========================//
func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)

	return nil
}

// ==============================//
// clean up func
// ==============================//
func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
