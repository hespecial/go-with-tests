package main

const repeatCount = 5

func Repeat(character string) (res string) {
	for i := 0; i < repeatCount; i++ {
		res += character
	}
	return
}
