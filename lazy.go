package fpgo

import "sync"

type lazyInt func() int

func createLazyInt(f func() int) lazyInt {
	var v int
	var once sync.Once

	return func() int {
		once.Do(func() {
			v = f()
		})

		return v
	}
}
