package cache

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
)

type Cache[K comparable, V any] struct {
	data        map[K]V
	fn          RefreshFunc[K, V]
	handleErrFn HandleErrFunc
	cron        *cron.Cron
}

func New[K comparable, V any](fn RefreshFunc[K, V], handleErrFn HandleErrFunc, options ...Option[K, V]) *Cache[K, V] {
	c := &Cache[K, V]{
		data:        make(map[K]V),
		fn:          fn,
		handleErrFn: handleErrFn,
		cron:        cron.New(),
	}
	for _, option := range options {
		option(c)
	}
	c.cron.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	go func() {
		select {
		case <-ch:
			fmt.Println("Stopping the cron job engine")
			c.cron.Stop()
			os.Exit(0)
		}
	}()

	return c
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache[K, V]) Refresh() {
	result, err := c.fn()
	if err != nil {
		c.handleErrFn(err)
		return
	}
	c.data = result
}
