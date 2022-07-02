package minmax

func MinMax(list []float64) (float64, float64) {
	min := list[0]
	max := list[0]
	for _, intnumber := range list {
		if intnumber > max {
			max = intnumber
		}
		if intnumber < min {
			min = intnumber
		}
	}
	return min, max
}

func MinMaxData(list []float64) []float64 {
	min, max := MinMax(list)
	for pos, elem := range list {
		list[pos] = (elem - min) / (max - min)
	}
	return list
}
