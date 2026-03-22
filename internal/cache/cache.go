package cache

import (
	"fmt"

	"github.com/MostafaSensei106/Simple-Cache/internal/constants"
	"github.com/MostafaSensei106/Simple-Cache/internal/models"
)

type Cache struct {
	Queue models.Queue
	Hash  models.Hash
}

func NewCache() Cache {
	return Cache{
		Queue: models.NewQueue(),
		Hash:  models.Hash{},
	}
}

func (c *Cache) Check(word string) {
	var node *models.Node

	if val, ok := c.Hash[word]; ok {
		node = c.Remove(val)
	} else {
		node = &models.Node{Value: word}
	}
	c.Add(node)
	c.Hash[word] = node
}

func (c *Cache) Remove(n *models.Node) *models.Node {
	fmt.Printf(" removeing node: %s", n.Value)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length -= 1
	delete(c.Hash, n.Value)

	return n
}

func (c *Cache) Add(n *models.Node) {
	fmt.Printf("add: %s", n.Value)

	temp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n

	c.Queue.Length++

	if c.Queue.Length > constants.MAX_SIZE {
		fmt.Printf(" max size reached")
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Display() {
	c.Queue.Display()
}
