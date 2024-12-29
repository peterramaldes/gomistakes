package client

import (
	"fmt"

	"github.com/peterramaldes/gomistakes/interfaces/store"
)

type Store interface{}

func Foo() {
	mem := store.NewInMemoryStore()
	fmt.Printf("mem: %v\n", mem)
}
