package simple

import "github.com/google/wire"

//============================//
//Binding Interface
//============================//

type SayHello interface {
	Hello(name string) string
}

// Ini bentuk anonymous field (embedded), artinya HelloService akan menyisipkan interface SayHello sebagai field.
// Struct ini otomatis mewarisi method dari field SayHello — bisa kamu panggil langsung HelloService.Hello("Nama").
type HelloService struct {
	SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

// provider
func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

// provider
func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{
		SayHello: sayHello,
	}
}

// SALAH
//func InitializedHelloService() *HelloService {
//	wire.Build(NewHelloService, NewSayHelloImpl)
//
//	return nil
//}

//cara manual
//var sayHello = NewSayHelloImpl()
//helloService := NewHelloService(sayHello)

// dengan wire
// membuat provider set
var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)
