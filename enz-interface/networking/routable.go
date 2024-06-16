package networking

type RoutableItem interface {
	Destinations() ([]string, error)
}
