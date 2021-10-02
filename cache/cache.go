package cache

import (
	"crypto/rand"
	"sync"
	"time"
)

const (
	letterBytes = "abcdefghipqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	maxTTL      = 24 * time.Hour
)

// https://www.socketloop.com/tutorials/golang-how-to-generate-random-string
func randString() string {
	var bytes = make([]byte, 16)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = letterBytes[v%byte(len(letterBytes))]
	}
	return string(bytes)
}

type entry struct {
	item   interface{}
	expire time.Time
}

type cache struct {
	ttl     time.Duration
	entries map[string]entry
	mux     sync.RWMutex
}

func New(ttl time.Duration) *cache {
	c := &cache{ttl: ttl, entries: map[string]entry{}}
	go c.monitor()
	return c
}

func (c *cache) Add(e interface{}) string {
	s := randString()
	c.mux.Lock()
	c.entries[s] = entry{item: e, expire: time.Now().Add(c.ttl)}
	c.mux.Unlock()
	return s
}

func (c *cache) AddWithTTL(e interface{}, ttl time.Duration) string {
	s := randString()
	if ttl > maxTTL {
		ttl = maxTTL
	}
	c.mux.Lock()
	c.entries[s] = entry{item: e, expire: time.Now().Add(ttl)}
	c.mux.Unlock()
	return s
}

func (c *cache) Get(s string) (interface{}, bool) {
	c.mux.RLock()
	v, ok := c.entries[s]
	c.mux.RUnlock()
	return v.item, ok
}

func (c *cache) Delete(s string) {
	c.mux.Lock()
	delete(c.entries, s)
	c.mux.Unlock()
}

func (c *cache) expireEntries() {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.entries {
		if v.expire.Before(time.Now().UTC()) {
			delete(c.entries, k)
		}
	}
}

func (c *cache) monitor() {
	for {
		c.expireEntries()
		time.Sleep(500 * time.Millisecond)
	}
}
