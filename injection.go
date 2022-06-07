package headwater

type Injector[T any] interface {
	Get() (bool, T)
}

type valueInjector[T any] struct {
	value T
}

func (i *valueInjector[T]) Get() (bool, T) {
	if i == nil {
		return false, GetZero[T]()
	}
	return true, i.value
}

func CreateValue[T any](value T) Injector[T] {
	return &valueInjector[T]{value}
}

type factory[T any] func() T

type factoryInjector[T any] struct {
	factory factory[T]
}

func (i *factoryInjector[T]) Get() (bool, T) {
	if i == nil {
		return false, GetZero[T]()
	}
	return true, i.factory()
}

func CreateFactory[T any](factory factory[T]) Injector[T] {
	return &factoryInjector[T]{factory}
}

type singletonInjector[T any] struct {
	factory factory[T]
	value   T
	ok      bool
	done    bool
}

func (i *singletonInjector[T]) Get() (bool, T) {
	if i == nil {
		return false, GetZero[T]()
	}
	if !i.done {
		i.done = true
		i.ok = true
		i.value = i.factory()
	}
	return i.ok, i.value
}

func CreateSingleton[T any](factory factory[T]) Injector[T] {
	return &singletonInjector[T]{
		factory: factory,
		ok:      false,
		done:    false,
	}
}
