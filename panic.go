package must

// Do alias for Do2
func Do[T any](result T, err error) T {
	return Do2(result, err)
}

// Do1 call panic if err not nil
func Do1(err error) {
	if err != nil {
		panic(err)
	}
}

// Do2 call panic if err not nil else return result
func Do2[T any](result T, err error) T {
	Do1(err)
	return result
}

// Do3 call panic if err not nil else return results
func Do3[T, R any](result1 T, result2 R, err error) (T, R) {
	Do1(err)
	return result1, result2
}

// Do4 call panic if err not nil else return results
func Do4[T, R, S any](result1 T, result2 R, result3 S, err error) (T, R, S) {
	Do1(err)
	return result1, result2, result3
}
