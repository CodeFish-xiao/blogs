package decorator

func GenericDecorator[T any](a T) (T, error) {

	return a, nil
}
