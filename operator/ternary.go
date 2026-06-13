package operator

// Idea from https://github.com/samber/lo

func Ternary[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func TernaryF[T any](condition bool, ifTrue func() T, ifFalse func() T) T {
	if condition {
		return ifTrue()
	}
	return ifFalse()
}
