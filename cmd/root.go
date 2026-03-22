package cmd

import (
	"fmt"

	"github.com/MostafaSensei106/Simple-Cache/internal/cache"
)

func Execute() {
	fmt.Println("Simple Cache Start...")

	c := cache.NewCache()
	for _, word := range []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi"} {
		c.Check(word)
		c.Display()
	}
	fmt.Println("Simple Cache End...")
}
