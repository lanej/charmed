// Package rendercache retains one string projection per component-owned cache.
package rendercache

import "strings"

// Cache retains only its current key and independently owned output. Callers
// serialize access and include every visible dependency in K or call Clear.
type Cache[K comparable] struct {
	key   K
	value string
	valid bool
}

func (c *Cache[K]) Get(key K, render func() string) string {
	if !c.valid || c.key != key {
		c.value = strings.Clone(render())
		c.key, c.valid = key, true
	}

	return c.value
}
func (c *Cache[K]) Clear() { var zero K; c.key = zero; c.value = ""; c.valid = false }
