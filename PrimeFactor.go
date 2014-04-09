package PrimeFactor

func PrimeFactor( number int ) ( factors []int ) {
	factors = []int{}
	
	for factor := 2; factor <= number; factor++ {
		for (number % factor == 0) {
			factors = append(factors, factor)
			number /= factor
		}
	}

	return
}