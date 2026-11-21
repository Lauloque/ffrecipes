package kitchen

type Recipe interface {
	Name() string
	Description() string
	Cook(*Pantry) error
}
