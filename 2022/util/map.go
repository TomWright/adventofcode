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

func ContainsCount[T comparable](item T, set []T) int {
	res := 0
	for _, v := range set {
		if v == item {
			res++
		}
	}
	return res
}

func Filter[T any](set []T, fn func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range set {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

func Reverse[T any](set []T) []T {
	res := make([]T, len(set))
	for i, j := 0, len(set)-1; i < len(set); i, j = i+1, j-1 {
		res[i] = set[j]
	}
	return res
}
