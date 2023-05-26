package slice

func Apply[T any](s []T, f func(T) T) []T {
	result := make([]T, len(s))

	for k, v := range s {
		result[k] = f(v)
	}

	return result
}

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue

	for _, v := range s {
		acc = f(acc, v)
	}

	return acc
}

func Filter[T any](s []T, f func(T) bool) []T {
	var result []T

	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}
