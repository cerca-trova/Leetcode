package feature

func IsPalindrome(x int) bool {
	var signal bool
	if x < 0 {
		signal = false
	}
	if x == 0 {
		signal = true
	}

	if x > 0 {
		ret := x
		y := 0
		d := calDigits(x)
		for i := d; i > 0; i-- {
			last_digit := x % 10
			x = (x - last_digit) / 10
			y = y + last_digit*micheng(i)
		}
		if y == ret {
			signal = true
		} else {
			signal = false
		}
	}
	return signal

}

func calDigits(x int) int {
	var digits int = 1
	for i := x; i >= 10; i /= 10 {
		digits++
	}
	return digits
}

func micheng(d int) int {
	s := 1
	for i := 1; i < d; i++ {
		s *= 10
	}
	return s
}
