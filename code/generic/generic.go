package generic

type Numeric interface {
	int | int64 | int32
}
type Numeric2 interface {
	float32 | float64
}
type FuckNumeric interface {
	Numeric2 | Numeric
}
type genericOption func()

func Add[T FuckNumeric](a, b T) T {
	return a + b
}
