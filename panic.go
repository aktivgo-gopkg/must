package unexpected

// Must call panic if err not nil else return result
func Must[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}

	return result
}

// Panic call panic if err not nil
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
