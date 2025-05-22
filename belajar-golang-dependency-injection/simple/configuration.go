package simple

// ==========================//
// Struct Field Provider
// ==========================//
type Configuration struct {
	Name string
}

type Application struct {
	configuration *Configuration
}

// provider
func NewApplication() *Application {
	return &Application{
		configuration: &Configuration{
			Name: "Rizky Cahyono Putra",
		},
	}
}
