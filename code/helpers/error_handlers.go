package helpers

func Check_error(e error) {
	if e != nil {
		panic(e)
	}
}
