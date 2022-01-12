package ignite

import "reflect"

// creates a new instance of type T
func New[T any]() (i T) {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Pointer {
		return reflect.New(t.Elem()).Interface().(T)
	}
	return i
}
