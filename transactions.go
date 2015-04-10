package transaction


func Average(numbers []float64) (average float64) {
	average = 0
	for i := 0; i < len(numbers); i++ {
		average += numbers[i]
	}
	return average/float64(len(numbers))
}
