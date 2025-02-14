package utils

func PassOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
