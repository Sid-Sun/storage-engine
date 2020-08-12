package store

// Store wraps CollectionInterface
type Store struct {
	CollectionInterface
}

// NewStore creates a new Store instance
func NewStore(instanceInterface CollectionInterface) Store {
	return Store{instanceInterface}
}
