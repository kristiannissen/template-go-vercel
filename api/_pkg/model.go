package pkg

type Model struct {
	Name string
}

func NewModel() *Model {
	return & Model{Name: "Kitty"}
}
