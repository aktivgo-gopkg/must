package must

// All causes panic if the last argument is an error and is not nil
func All(args ...any) {
	if err, ok := args[len(args)-1].(error); ok && err != nil {
		panic(err)
	}
}

// Do causes panic if the error is not nil otherwise returns the result
func Do[T any](result T, err error) T {
	All(err)
	return result
}
