package iteration

//WRITE THE MINIMAL AMOUNT OF CODE FOR TEST TO RUN

func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ { //loop does not require paranthesis
		repeated += character
	}
	return repeated
}

func Compare(x, y string) bool {
	return x == y
}
