package homework

func average(input [15]float32) (result float32) {
	sum := 0
	num := 0
	for i := 0; i < len(input); i++ {
		sum += int(input[i])
		if int(input[i]) != 0 {
			num++
		}
	}
	result = float32(sum) / float32(num)
	return
}
