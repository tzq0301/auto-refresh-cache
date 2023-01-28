package cache

type RefreshFunc[K comparable, V any] func() (map[K]V, error)
