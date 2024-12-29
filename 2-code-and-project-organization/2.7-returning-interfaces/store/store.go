package store

type InMemory struct {
}

func NewInMemoryStore() InMemory {
	return InMemory{}
}
