package main

func drawRhombus(elements int) {
	//validating that it has n odd numbers
	if elements%2 == 0 {
		elements++
	}

	//Upper half of the rhombus
	for i := 1; i <= elements; i += 2 {
		spaces := (elements - i) / 2
		for j := 0; j < spaces; j++ {
			print(" ")
		}
		for j := 0; j < i; j++ {
			print("*")
		}
		println()
	}

	//lower half of rhombus
	for i := elements - 2; i >= 1; i -= 2 {
		spaces := (elements - i) / 2
		for j := 0; j < spaces; j++ {
			print(" ")
		}
		for j := 0; j < i; j++ {
			print("*")
		}
		println()
	}
}

func main() {
	drawRhombus(5)
}
