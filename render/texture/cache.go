package texture

import (
	"fmt"
)

type Cache struct {
	cache map[string]uint32
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]uint32),
	}
}

func (c *Cache) Load(name, path string) {
	c.cache[name] = FromFile(path)
}

func (c *Cache) Get(name string) uint32 {
	tex, ok := c.cache[name]
	if !ok {
		panic(fmt.Errorf("no such texture %q loaded", name))
	}
	return tex
}
