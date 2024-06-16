package networking

// Mesh is an interface that defines the methods that a Enzo-compatible mesh network must implement.
type MeshNetwork interface {
	// Enroll is a method that is used to add a new node to the mesh network, with all needed information.
	Enroll(node MeshNode) error
	// Unenroll is a method that is used to remove a node from the mesh network.
	// This method should be called when a node is no longer available.
	Unenroll(nodeId string) error
	// Transfer is a method that sends data to a specific node in the mesh network.
	Transfer(nodeId string, data RoutableItem) error
}
