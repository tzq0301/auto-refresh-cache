package cache

type Option[K comparable, V any] func(cache *Cache[K, V])

func RefreshOnStart[K comparable, V any](c *Cache[K, V]) {
	c.Refresh()
}

// RefreshWithCron refers to https://pkg.go.dev/github.com/robfig/cron
func RefreshWithCron[K comparable, V any](cronExpr string) Option[K, V] {
	return func(c *Cache[K, V]) {
		_, err := c.cron.AddFunc(cronExpr, func() {
			c.Refresh()
		})
		if err != nil {
			c.handleErrFn(err)
		}
	}
}
