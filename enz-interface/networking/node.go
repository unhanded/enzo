package networking

// MeshNode is an interface that defines the methods that a Enzo-compatible mesh node must implement.
type MeshNode interface {
	// GetId is a method that returns the unique identifier of the node.
	// This identifier must be unique in the mesh network.
	Id() string
	// Send is a method that sends data to a specific node in the mesh network.
	Recieve(data RoutableItem) error
}
