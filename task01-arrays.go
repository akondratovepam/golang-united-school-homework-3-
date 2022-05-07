package homework

func average(input [15]float32) (result float32) {
	len := float32(len(input))
	for i := range input {
		result += input[i] / len
	}
	return
}
