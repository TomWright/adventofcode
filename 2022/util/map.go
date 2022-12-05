package util

func Map[I any, O any](in []I, mapper func(I) O) []O {
	res := make([]O, len(in))
	for k, v := range in {
		res[k] = mapper(v)
	}
	return res
}

func First[T any](n int, in []T) []T {
	res := make([]T, n)
	for x := 0; x < n; x++ {
		res[x] = in[x]
	}
	return res
}

func Sum[T int | int32 | int64 | float32 | float64](in []T) T {
	var res T
	for _, v := range in {
		res += v
	}
	return res
}

func Contains[T comparable](item T, set []T) bool {
	for _, v := range set {
		if v == item {
			return true
		}
	}
	return false
}
