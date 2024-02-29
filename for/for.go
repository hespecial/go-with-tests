package main

func Repeat(character string, count int) (res string) {
	for i := 0; i < count; i++ {
		res += character
	}
	return
}
