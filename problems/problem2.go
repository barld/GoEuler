package problems

func GetAnswer2() int {
	sum := 0
	for f1, f2 := 1, 2; f2 < 4000000; {
		if f2%2 == 0 {
			sum += f2
		}
		f2, f1 = f2+f1, f2
	}
	return sum
}
